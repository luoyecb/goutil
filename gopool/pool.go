package gopool

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrPoolClosed = errors.New("Pool closed")
	ErrTimeout    = errors.New("Submit timeout")
	ErrNoCallback = errors.New("No callback")
)

type Pool struct {
	runChan  chan *Job
	callback func(interface{})
	workers  []*worker

	closed  bool
	closeMu sync.Mutex
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

func (p *Pool) Close() {
	p.closeMu.Lock()
	if p.closed {
		p.closeMu.Unlock()
		return
	}
	p.closed = true
	p.closeMu.Unlock()

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
