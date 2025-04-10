package strjoiner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringJoiner(t *testing.T) {
	assert := assert.New(t)

	sj := NewStringJoiner(",")
	assert.Equal(sj.Len(), 0)

	sj.Add("1")
	sj.Add("2")
	sj.Add("3")
	assert.Equal(sj.String(), "1,2,3")
	assert.Equal(sj.Len(), 5)
}

func TestStringJoiner3(t *testing.T) {
	assert := assert.New(t)

	sj := NewStringJoiner3(", ", "[", "]")
	assert.Equal(sj.Len(), 0)

	sj.Add("1")
	sj.Add("2")
	sj.Add("3")
	assert.Equal(sj.String(), "[1, 2, 3]")
	assert.Equal(sj.Len(), 9)
}

func TestStringJoiner5(t *testing.T) {
	assert := assert.New(t)

	sj := NewStringJoiner5(", ", "[", "]", "\"", "\"")
	assert.Equal(sj.Len(), 0)

	sj.Add("1").Add("2").Add("3")
	assert.Equal(sj.String(), `["1", "2", "3"]`)
	assert.Equal(sj.Len(), 15)
}
