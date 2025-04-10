package goutil

import (
	"bufio"
	"os"
	"strings"
)

type FileIter struct {
	file          *os.File
	scanner       *bufio.Scanner
	commentPrefix []string
}

func NewFileIter(filepath string, prefix ...string) (*FileIter, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return &FileIter{
		file:          f,
		scanner:       bufio.NewScanner(f),
		commentPrefix: prefix,
	}, nil
}

func NewStdinIter(prefix ...string) *FileIter {
	return &FileIter{
		scanner:       bufio.NewScanner(os.Stdin),
		commentPrefix: prefix,
	}
}

func NewStringIter(content string, prefix ...string) *FileIter {
	return &FileIter{
		scanner:       bufio.NewScanner(strings.NewReader(content)),
		commentPrefix: prefix,
	}
}

func (iter *FileIter) EachLine(fn func(string)) {
	if fn != nil {
		for iter.HasNext() {
			line := iter.Text()
			if line != "" {
				fn(line)
			}
		}
	}
}

func (iter *FileIter) HasNext() bool {
	return iter.scanner.Scan()
}

func (iter *FileIter) Text() string {
	for {
		text := iter.scanner.Text()
		if !iter.isComment(text) {
			return text
		}
		hasNext := iter.scanner.Scan()
		if !hasNext {
			break
		}
	}
	return ""
}

func (iter *FileIter) isComment(line string) bool {
	for _, prefix := range iter.commentPrefix {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func (iter *FileIter) Close() {
	if iter.file != nil {
		iter.file.Close()
	}
}
