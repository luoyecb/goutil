package goutil

import (
	"context"
	"sync"
)

func NewGo(fn func(context.Context)) context.CancelFunc {
	if fn == nil {
		return nil
	}
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fn(ctx)
	}()

	return func() {
		cancel()
		wg.Wait()
	}
}
