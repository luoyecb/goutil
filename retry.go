package goutil

import (
	"errors"
	"reflect"
)

type FuncRetry struct {
	value    reflect.Value
	numIn    int
	retries  int
	errIndex int
}

func NewFuncRetry(v interface{}, retries, errIndex int) *FuncRetry {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Func {
		panic("param must be func")
	}
	rt := rv.Type()
	if errIndex >= rt.NumOut() {
		panic("errIndex invalid")
	}

	return &FuncRetry{
		value:    rv,
		numIn:    rt.NumIn(),
		retries:  retries,
		errIndex: errIndex,
	}
}

func (r *FuncRetry) Call(args ...interface{}) ([]interface{}, error) {
	if len(args) != r.numIn {
		panic("number of param not equal")
	}

	input := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		input = append(input, reflect.ValueOf(arg))
	}

	var lastErr error
	callCnt := 0
	for callCnt < r.retries {
		output := r.value.Call(input)
		errIn := output[r.errIndex].Interface()
		if errIn == nil {
			results := make([]interface{}, 0, len(output))
			for _, res := range output {
				results = append(results, res.Interface())
			}
			return results, nil
		} else {
			lastErr = errIn.(error)
		}
		callCnt++
	}

	if lastErr != nil {
		return nil, lastErr
	}
	return nil, errors.New("too many retry")
}
