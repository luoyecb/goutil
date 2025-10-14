package gopool

import (
	"fmt"
)

type worker struct {
	runChan chan *Job // receive job

	closeCh chan struct{}
	stopCh  chan struct{}
}

func newWorker(ch chan *Job) *worker {
	w := &worker{
		runChan: ch,
		closeCh: make(chan struct{}),
		stopCh:  make(chan struct{}),
	}
	go w.run()
	return w
}

func (w *worker) run() {
	defer func() {
		close(w.stopCh)
	}()

	for {
		select {
		case job, ok := <-w.runChan:
			if !ok {
				return // exit
			}
			job.notify(callRunner(job.r))
		case <-w.closeCh:
			return // exit
		}
	}
}

func callRunner(r Runner) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered! err: %v\n", r)
		}
	}()
	r.Run()
	return
}

func (w *worker) close() {
	close(w.closeCh)
}

func (w *worker) wait() {
	<-w.stopCh
}
