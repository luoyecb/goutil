package assert

import (
	"fmt"
	"reflect"
)

var (
	assertStat = &AssertStat{}
)

type AssertStat struct {
	failed  int
	succeed int
}

func (s *AssertStat) String() string {
	return fmt.Sprintf("succeed: %d, failed: %d, total: %d", s.succeed, s.failed, s.succeed+s.failed)
}

func Stat() {
	fmt.Println(assertStat)
}

func Nil(v interface{}) {
	if v == nil || reflect.ValueOf(v).IsNil() {
		assertStat.succeed++
	} else {
		fmt.Printf("Nil assert failed, actual value: %v\n", v)
		assertStat.failed++
	}
}

func True(v bool) {
	if !v {
		fmt.Printf("True assert failed, actual value: %v\n", v)
		assertStat.failed++
	} else {
		assertStat.succeed++
	}
}

func False(v bool) {
	if v {
		fmt.Printf("False assert failed, actual value: %v\n", v)
		assertStat.failed++
	} else {
		assertStat.succeed++
	}
}

func Equal(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		fmt.Printf("Equal assert failed, actual value: a=%v, b=%v\n", a, b)
		assertStat.failed++
	} else {
		assertStat.succeed++
	}
}
