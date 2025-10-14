package gopool

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
