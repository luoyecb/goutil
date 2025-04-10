package goutil

import (
	"sync"
	"sync/atomic"
)

// 自实现sync.Once
type Once2 struct {
	done uint32
	mu   sync.Mutex
}

func (o *Once2) Do(fn func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	if o.done == 0 {
		fn()
		atomic.CompareAndSwapUint32(&o.done, 0, 1)
	}
}
