package goutil

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// test:TestConcurrent
func TestConcurrent(t *testing.T) {
	assert := assert.New(t)

	futures := Concurrent(3, func(_ int, input interface{}) (interface{}, error) {
		s := input.(string)
		fmt.Println(s)

		time.Sleep(time.Second)
		return s, nil
	}, "hello", "world", "golang", "java")

	for _, f := range futures {
		f.Wait()
		fmt.Printf("result: %s, index: %d, err: %v\n", f.Result(), f.Index(), f.Err())
	}

	assert.NotNil(futures)
}
