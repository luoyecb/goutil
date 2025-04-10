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
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		totalChars:     100,
		style:          NewStyle("#", "", " "),
		tickerDuration: time.Second,
	}
}

func NewProgressBar2() *ProgressBar {
	return &ProgressBar{
		totalChars:     100,
		style:          NewStyle("=", ">", "-"),
		tickerDuration: time.Second,
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
	b.start = time.Now()
	b.ticker = time.NewTicker(b.tickerDuration)
	go func() {
		for {
			select {
			case <-b.ticker.C:
				b.printProgress()
			}
		}
	}()
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
	if b.ticker != nil {
		b.ticker.Stop()
	}
	b.printProgress()
	fmt.Println()
}
