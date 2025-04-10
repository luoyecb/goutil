package urand

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandRange(t *testing.T) {
	var tests = []struct {
		min int64
		max int64
	}{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 10},
		{11, 20},
		{21, 30},
		{31, 100},
		{101, 1000},
		{1001, 10000},
	}

	for _, test := range tests {
		v := RandRange(test.min, test.max)
		fmt.Printf("RandRange(%[1]d, %[2]d) = %d\n", test.min, test.max, v)
		if v < test.min || v > test.max {
			t.Errorf("RandRange(%d, %d) = %d, not in [%[1]d, %[2]d]", test.min, test.max, v)
		}
	}
}

func TestRandString(t *testing.T) {
	var min = 4
	var max = 10
	for i, j := 0, 10000; i < j; i++ {
		str := RandString(min, max)
		assert.True(t, len(str) >= min)
		assert.True(t, len(str) <= max)
		// fmt.Println(str)
	}
}
