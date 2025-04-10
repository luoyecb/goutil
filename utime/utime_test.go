package utime

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	// assert := assert.New(t)
	fmt.Println(Unixtime())
	fmt.Println(UnixNano())
	fmt.Println(Date())
	fmt.Println(Time())
	fmt.Println(Datetime())
}

func TestSleep(t *testing.T) {
	assert := assert.New(t)

	now := Unixtime()
	Sleep(3)
	assert.Equal(now+3, Unixtime())
}

func TestTick(t *testing.T) {
	assert := assert.New(t)

	now := Unixtime()
	cancel := Tick(1, func() {
		fmt.Println(Unixtime())
		now++
	})

	// 3秒后取消
	After(3, func() {
		cancel()
	})
	assert.Equal(now, Unixtime())
}
