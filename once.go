package goutil

// 自实现sync.Once
type Once struct {
	ch chan struct{}
}

func NewOnce() *Once {
	o := &Once{
		ch: make(chan struct{}, 1),
	}
	o.ch <- struct{}{}
	return o
}

func (o *Once) Do(f func()) {
	// 阻塞在此处的goroutine需要等待
	if _, ok := <-o.ch; !ok {
		return
	}

	f()
	close(o.ch)
}
