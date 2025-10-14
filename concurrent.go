package goutil

import (
	"fmt"
)

type FutureFunc func(index int, input interface{}) (result interface{}, err error)

type Future struct {
	index int // input index
	input interface{}

	result interface{}
	err    error // not nil, means that error has occurred

	done chan struct{} // task completed?
}

func NewFuture() *Future {
	return &Future{
		done: make(chan struct{}, 1),
	}
}

func (f *Future) Index() int {
	return f.index
}

func (f *Future) Input() interface{} {
	return f.input
}

func (f *Future) Err() error {
	return f.err
}

func (f *Future) Result() interface{} {
	return f.result
}

// Wait for task completed
func (f *Future) Wait() {
	<-f.done
}

func Concurrent(conc int, fn FutureFunc, inputs ...interface{}) []*Future {
	if len(inputs) < 1 || conc < 1 {
		return nil
	}
	if conc > len(inputs) {
		conc = len(inputs)
	}

	tasks := make(chan *Future, len(inputs))

	// make Future
	futures := make([]*Future, 0, len(inputs))
	for index, input := range inputs {
		f := NewFuture()
		f.index = index
		f.input = input

		futures = append(futures, f)
		tasks <- f
	}
	close(tasks)

	// make goroutines
	for i := 0; i < conc; i++ {
		go func() {
			for {
				select {
				case f, ok := <-tasks:
					if !ok {
						// fmt.Println("Exit")
						return // exit
					}
					// exec
					runFuture(f, fn)
				}
			}
		}()
	}

	return futures
}

func runFuture(f *Future, fn FutureFunc) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				f.err = err
			} else {
				f.err = fmt.Errorf("panic recovered: %v", r)
			}
			f.result = nil
		}

		f.done <- struct{}{}
	}()

	f.result, f.err = fn(f.index, f.input)
}
