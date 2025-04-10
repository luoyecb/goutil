package alignment

import (
	"bufio"
	"bytes"
	"strings"
)

type Alignment struct {
	sep string
}

func NewAlignment(sep string) *Alignment {
	return &Alignment{sep}
}

func (a *Alignment) Format(text string) string {
	if text == "" {
		return ""
	}

	var buf bytes.Buffer
	lines := []string{}

	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
			continue
		}
		// 按照空行分隔为多块处理
		if len(lines) > 0 {
			buf.WriteString(a.format(lines))
			buf.WriteByte('\n')
			lines = []string{}
		} else {
			buf.WriteByte('\n')
		}
	}
	if len(lines) > 0 {
		buf.WriteString(a.format(lines))
	}
	if scanner.Err() != nil {
		return text
	}

	return buf.String()
}

func (a *Alignment) format(lines []string) string {
	maxP1Len := 0
	sepLen := len(a.sep)
	content := make([]*Line, 0, len(lines))

	for _, line := range lines {
		var aLine *Line
		var p1 string

		if index := strings.Index(line, a.sep); index == -1 {
			p1 = strings.TrimSpace(line)
			aLine = NewLine(p1, "", false, a)
		} else {
			p1 = strings.TrimSpace(line[:index])
			aLine = NewLine(p1, strings.TrimSpace(line[index+sepLen:]), true, a)
		}
		if len(p1) > maxP1Len {
			maxP1Len = len(p1)
		}

		content = append(content, aLine)
	}

	var buf bytes.Buffer
	for _, l := range content {
		buf.WriteString(l.format(maxP1Len))
		buf.WriteByte('\n')
	}
	return buf.String()
}

type Line struct {
	p1     string
	p2     string
	hasSep bool
	align  *Alignment
}

func NewLine(p1, p2 string, b bool, a *Alignment) *Line {
	return &Line{p1, p2, b, a}
}

func (l *Line) format(maxsize int) string {
	if !l.hasSep {
		return l.p1
	}

	var buf bytes.Buffer

	buf.WriteString(l.p1)
	for i := maxsize - len(l.p1); i >= 0; i-- {
		buf.WriteByte(' ')
	}
	buf.WriteString(l.align.sep)
	buf.WriteByte(' ')
	buf.WriteString(l.p2)

	return buf.String()
}
