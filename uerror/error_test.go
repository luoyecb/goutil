package uerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dummyErr = errors.New("dummy error")
)

func TestCheckErr(t *testing.T) {
	assert := assert.New(t)

	CheckErr(dummyErr, func(err error) {
		assert.True(err == dummyErr)
	})
}

func TestWithRecover(t *testing.T) {
	assert := assert.New(t)

	err := WithRecover(func() error {
		panic(dummyErr)
		return nil
	}, func(err error) {
		assert.True(err == dummyErr)
	})
	assert.True(err == dummyErr)
}

func TestWithRecover2(t *testing.T) {
	assert := assert.New(t)

	err := WithRecover(func() error {
		return dummyErr
	}, func(err error) {
		assert.True(err == dummyErr)
	})
	assert.True(err == dummyErr)
}
