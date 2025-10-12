package ufile

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUfile(t *testing.T) {
	assert := assert.New(t)

	goroot := runtime.GOROOT()
	fmt.Println(goroot)
	assert.True(FileExists(goroot))
	assert.False(FileExists("FileExists(goroot)"))
}
