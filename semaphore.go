package goutil

// 信号量
type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(cnt int) *Semaphore {
	if cnt <= 0 {
		return nil
	}
	return &Semaphore{
		ch: make(chan struct{}, cnt),
	}
}

func (s *Semaphore) Accquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}
