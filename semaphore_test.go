package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test:TestSemaphore
func TestSemaphore(t *testing.T) {
	assert := assert.New(t)

	sum := 0
	cnt := 10000
	sem := NewSemaphore(1)

	concurrent := 10
	ch := ConcurrentRun(func() {
		for i := 0; i < cnt; i++ {
			sem.Accquire()
			sum++
			sem.Release()
		}
	}, concurrent)
	<-ch

	assert.True(sum == cnt*concurrent)
}
