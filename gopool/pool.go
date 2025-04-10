package gopool

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrPoolClosed = errors.New("Pool closed")
	ErrTimeout    = errors.New("Submit timeout")
	ErrNoCallback = errors.New("No callback")
)

// ========================================
// Runner
// ========================================
type Runner interface {
	Run()
}

type PoolFunc func()

func (pf PoolFunc) Run() {
	pf()
}

type Callback struct {
	input interface{}
	fn    func(interface{})
}

func (c *Callback) Run() {
	c.fn(c.input)
}

// ========================================
// Job
// ========================================
type Job struct {
	r      Runner
	err    error
	doneCh chan struct{}
}

func NewJob(r Runner) *Job {
	return &Job{
		r:      r,
		doneCh: make(chan struct{}),
	}
}

func (j *Job) notify(err error) {
	j.err = err
	close(j.doneCh)
}

func (j *Job) Err() error {
	return j.err
}

func (j *Job) Wait() {
	<-j.doneCh
}

// ========================================
// worker
// ========================================
type worker struct {
	runChan chan *Job
	stop    chan struct{}
	wg      sync.WaitGroup
}

func newWorker(ch chan *Job) *worker {
	w := &worker{
		runChan: ch,
		stop:    make(chan struct{}),
	}
	w.wg.Add(1)
	go w.run()
	return w
}

func (w *worker) run() {
	defer w.wg.Done()

	for {
		select {
		case job, ok := <-w.runChan:
			if !ok {
				return
			}
			job.notify(callRunner(job.r))
		case <-w.stop:
			return // exit
		}
	}
}

func callRunner(r Runner) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered! err: %v\n", err)
		}
	}()
	r.Run()
	return
}

func (w *worker) close() {
	close(w.stop)
}

func (w *worker) wait() {
	w.wg.Wait()
}

// ========================================
// Pool
// ========================================
type Pool struct {
	runChan  chan *Job
	callback func(interface{})
	workers  []*worker

	closed bool
}

func NewPool(maxGoroutines, queueCap int) *Pool {
	return NewPoolCallback(maxGoroutines, queueCap, nil)
}

func NewPoolCallback(maxGoroutines, queueCap int, cb func(interface{})) *Pool {
	p := &Pool{
		runChan:  make(chan *Job, queueCap),
		callback: cb,
	}
	for i := 0; i < maxGoroutines; i++ {
		p.workers = append(p.workers, newWorker(p.runChan))
	}
	return p
}

func (p *Pool) Close() {
	if p.closed {
		return
	}
	p.closed = true

	for _, w := range p.workers {
		w.close()
	}

	close(p.runChan)
	for job := range p.runChan {
		job.notify(ErrPoolClosed)
	}

	for _, w := range p.workers {
		w.wait()
	}
}

func (p *Pool) SubmitFunc(fn func()) (*Job, error) {
	return p.Submit(PoolFunc(fn))
}

func (p *Pool) SubmitInput(input interface{}) (*Job, error) {
	if p.callback == nil {
		return nil, ErrNoCallback
	}
	return p.Submit(&Callback{input: input, fn: p.callback})
}

func (p *Pool) Submit(r Runner) (*Job, error) {
	return p.SubmitTimeout(r, -1)
}

func (p *Pool) SubmitTimeout(r Runner, timeoutMilli int) (job *Job, err error) {
	defer func() {
		if r := recover(); r != nil {
			job = nil
			err = ErrPoolClosed
		}
	}()

	job = NewJob(r)
	if timeoutMilli <= 0 {
		p.runChan <- job
	} else {
		select {
		case p.runChan <- job:
		case <-time.After(time.Duration(timeoutMilli) * time.Millisecond):
			return nil, ErrTimeout
		}
	}
	return job, nil
}
