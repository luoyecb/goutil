package gopool

type Runner interface {
	Run()
}

type PoolFunc func()

func (pf PoolFunc) Run() {
	pf()
}

type Callback struct {
	input interface{}
	fn    func(interface{})
}

func (c *Callback) Run() {
	c.fn(c.input)
}
