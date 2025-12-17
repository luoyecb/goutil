package goutil

import (
	"sync"
)

func ConcurrentRun(fn func(), concurrent int) chan struct{} {
	stop := make(chan struct{})
	if concurrent < 1 {
		close(stop)
		return stop
	}

	var wg sync.WaitGroup
	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}

	wg.Wait()
	close(stop)
	return stop
}
