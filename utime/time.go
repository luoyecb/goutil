package utime

import (
	"time"
)

const (
	DATE_STR     = "2006-01-02"
	TIME_STR     = "15:04:05"
	DATETIME_STR = "2006-01-02 15:04:05"
)

func Unixtime() int64 {
	return time.Now().Unix()
}

func UnixNano() int64 {
	return time.Now().UnixNano()
}

func Now() string {
	return Datetime()
}

func Date() string {
	return time.Now().Format(DATE_STR)
}

func Time() string {
	return time.Now().Format(TIME_STR)
}

func Datetime() string {
	return time.Now().Format(DATETIME_STR)
}
