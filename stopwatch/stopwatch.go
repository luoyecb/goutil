package stopwatch

import (
	"fmt"
	"time"
)

type StopWatch struct {
	start int64
}

func NewStopWatch() *StopWatch {
	return &StopWatch{time.Now().UnixNano()}
}

// Return millisecond
func (st *StopWatch) ElapsedTime() int64 {
	now := time.Now().UnixNano()
	return (now - st.start) / 1000000
}

func (st *StopWatch) ElapsedTimeString() string {
	return st.String()
}

func (st *StopWatch) String() string {
	return fmt.Sprintf("Elapsed time: %d milliseconds.", st.ElapsedTime())
}
