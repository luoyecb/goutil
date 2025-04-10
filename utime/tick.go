package utime

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func Tick(seconds int, handleFunc func()) func() {
	if handleFunc == nil {
		return nil
	}

	var wg sync.WaitGroup
	stop := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(time.Duration(seconds) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-stop:
				return
			case <-ticker.C:
				withRecover(handleFunc)
			}
		}
	}()

	return func() {
		close(stop)
		wg.Wait()
	}
}

func After(seconds int, handleFunc func()) {
	if handleFunc == nil {
		return
	}

	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		withRecover(handleFunc)
	}
}

func withRecover(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "withRecover catch err: %+v\n", r)
		}
	}()
	fn()
}
