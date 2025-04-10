package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemaphore(t *testing.T) {
	assert := assert.New(t)

	sem := NewSemaphore(1)

	sum := 0
	cnt := 10000
	concurrent := 10

	ch := ConcurrentRun(func() error {
		for i := 0; i < cnt; i++ {
			sem.Accquire()
			sum++
			sem.Release()
		}
		return nil
	}, concurrent)
	<-ch

	assert.True(sum == cnt*concurrent)
}
