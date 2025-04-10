package uerror

import (
	"fmt"
	"log"
	"os"
)

func CheckErr(err error, handlers ...func(error)) {
	if err != nil {
		if len(handlers) > 0 {
			handlers[0](err)
		} else {
			panic(err)
		}
	}
}

func ErrExit(err error) {
	CheckErr(err, func(err error) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	})
}

func WithRecover(fn func() error, handlers ...func(error)) (err error) {
	if fn == nil {
		return
	}

	var errFn func(error)
	if len(handlers) > 0 {
		errFn = handlers[0]
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("WithRecover caught error: %+v\n", r)
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("WithRecover caught error: %+v\n", r)
			}
			if errFn != nil {
				errFn(err)
			}
		}
	}()

	err = fn()
	if err != nil && errFn != nil {
		errFn(err)
	}
	return
}
