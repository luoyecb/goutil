package gopool

import (
	"fmt"
	"log"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

var idCnt int64

func GetId() int64 {
	return atomic.AddInt64(&idCnt, 1)
}

type namePrinter struct {
	name string
}

func (np *namePrinter) Run() {
	time.Sleep(time.Second * 3)
	log.Println(np.name)
}

func newNamePrinter() *namePrinter {
	return &namePrinter{fmt.Sprintf("printer-%d", GetId())}
}

func TestPool(t *testing.T) {
	taskSize := 10
	pool := NewPool(3, taskSize)
	defer pool.Close()

	for i := 0; i < taskSize; i++ {
		log.Printf("Submit %d\n", i)
		_, err := pool.Submit(newNamePrinter())
		if err != nil {
			log.Println(err)
		}
	}

	for i := 0; i < taskSize; i++ {
		log.Printf("Submit %d\n", taskSize+i)
		_, err := pool.SubmitTimeout(newNamePrinter(), 1000)
		if err != nil {
			log.Println(err)
		}
		if i == 5 {
			pool.Close()
			fmt.Println("closed")
		}
	}
}

func TestPool2(t *testing.T) {
	defer func() {
		fmt.Printf("end goroutine num: %d\n", runtime.NumGoroutine())
	}()

	pool := NewPool(10, 10)
	defer pool.Close()

	fmt.Printf("start goroutine num: %d\n", runtime.NumGoroutine())

	pool.Submit(PoolFunc(func() {
		for i := 0; i < 20; i++ {
			fmt.Printf("goroutine num: %d\n", runtime.NumGoroutine())
			time.Sleep(time.Duration(500) * time.Millisecond)
		}
	}))
	pool.Submit(PoolFunc(func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("call runtime.Goexit")
		runtime.Goexit()
	}))
	pool.Submit(PoolFunc(func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("call panic")
		panic("Test FixedPool")
	}))

	fmt.Println("wait.")
	time.Sleep(time.Duration(15) * time.Second)
}
