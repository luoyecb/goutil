package goutil

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Style struct {
	LeftGraph  string
	MidGraph   string
	RightGraph string

	color *color.Color
}

func NewStyle(left, middle, right string) *Style {
	return &Style{left, middle, right, color.New(color.BgGreen)}
}

func (s *Style) String(leftCnt, rightCnt int) string {
	if s.MidGraph != "" {
		if leftCnt > 0 {
			leftCnt--
		} else {
			rightCnt--
		}
	}
	// return s.color.Sprintf("%s%s%s", strings.Repeat(s.LeftGraph, leftCnt), s.MidGraph, strings.Repeat(s.RightGraph, rightCnt))
	return strings.Repeat(s.LeftGraph, leftCnt) + s.MidGraph + strings.Repeat(s.RightGraph, rightCnt)
}

type ProgressBar struct {
	percent int
	mu      sync.Mutex

	totalChars int
	style      *Style

	start          time.Time
	ticker         *time.Ticker
	tickerDuration time.Duration
	stop           chan bool // receive stop signal
}

func NewProgressBar() *ProgressBar {
	return NewProgressBarStyle(NewStyle("#", "", " "))
}

func NewProgressBar2() *ProgressBar {
	return NewProgressBarStyle(NewStyle("=", ">", "-"))
}

func NewProgressBarStyle(st *Style) *ProgressBar {
	return &ProgressBar{
		totalChars:     100,
		style:          st,
		tickerDuration: time.Second,
		// stop:           make(chan bool),
	}
}

func (b *ProgressBar) Report(percent int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if percent > 100 {
		percent = 100
	} else if percent < 0 {
		percent = 0
	}
	b.percent = percent
}

func (b *ProgressBar) Start() {
	b.reset()

	if b.ticker == nil {
		b.ticker = time.NewTicker(b.tickerDuration)
	} else {
		b.ticker.Reset(b.tickerDuration)
	}
	go func() {
		for {
			select {
			case <-b.ticker.C:
				b.printProgress()
			case <-b.stop:
				return
			}
		}
	}()
}

func (b *ProgressBar) reset() {
	b.mu.Lock()
	b.percent = 0
	b.mu.Unlock()

	b.stop = make(chan bool)
	b.start = time.Now()
}

func (b *ProgressBar) printProgress() {
	b.mu.Lock()
	tmpPercent := b.percent
	b.mu.Unlock()

	leftCnt := b.totalChars * tmpPercent / 100
	rightCnt := b.totalChars - leftCnt
	cost := time.Since(b.start).Round(time.Millisecond)
	fmt.Printf("\r[%s]\t%d%% %10s", b.style.String(leftCnt, rightCnt), tmpPercent, cost)
}

func (b *ProgressBar) Stop() {
	if b.stop == nil {
		return // no start
	}
	_, ok := <-b.stop
	if !ok {
		return // has closed
	}
	close(b.stop)

	if b.ticker != nil {
		b.ticker.Stop()
		b.ticker = nil
	}
	b.printProgress()
	fmt.Println()
}
