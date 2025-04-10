package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce(t *testing.T) {
	assert := assert.New(t)

	once := NewOnce()
	cnt := 0

	ch := ConcurrentRun(func() error {
		once.Do(func() {
			cnt++
		})

		return nil
	}, 10000)
	<-ch

	assert.True(cnt == 1)
}

func TestOnce2(t *testing.T) {
	assert := assert.New(t)

	var once Once2
	cnt := 0

	ch := ConcurrentRun(func() error {
		once.Do(func() {
			cnt++
		})

		return nil
	}, 10000)
	<-ch

	assert.True(cnt == 1)
}
