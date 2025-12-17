package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test:TestOnce
func TestOnce(t *testing.T) {
	assert := assert.New(t)

	cnt := 0
	once := NewOnce()

	ch := ConcurrentRun(func() {
		once.Do(func() {
			cnt++
		})
	}, 10000)
	<-ch

	assert.True(cnt == 1)
}

func TestOnce2(t *testing.T) {
	assert := assert.New(t)

	cnt := 0
	var once Once2

	ch := ConcurrentRun(func() {
		once.Do(func() {
			cnt++
		})
	}, 10000)
	<-ch

	assert.True(cnt == 1)
}
