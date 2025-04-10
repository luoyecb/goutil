package goutil

import (
	"os"
	"os/signal"
)

// Ctrl+C
func Interrupt(fn func()) {
	Signal(fn, os.Interrupt)
}

func Signal(fn func(), sig ...os.Signal) {
	if len(sig) <= 0 {
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)

	go func() {
		<-sigCh
		fn()
	}()
}
