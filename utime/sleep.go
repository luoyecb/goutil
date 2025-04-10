package utime

import (
	"time"
)

func Sleep(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

func Msleep(millisec int) {
	time.Sleep(time.Duration(millisec) * time.Millisecond)
}
