package template

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

type Template struct {
	begin byte
	end   byte
	chars []byte
	size  int
	pos   int
}

func NewTemplate(text string) *Template {
	return NewTemplate3(text, '{', '}')
}

func NewTemplate3(text string, begin, end byte) *Template {
	t := &Template{
		begin: begin,
		end:   end,
		chars: []byte(text),
	}
	t.size = len(t.chars)
	return t
}

func (t *Template) nextChar(isPeek bool) (byte, bool) {
	if t.pos < t.size {
		ch := t.chars[t.pos]
		if !isPeek {
			t.pos++
		}
		return ch, true
	}
	return 0, false
}

func (t *Template) Reset() {
	t.pos = 0
}

func (t *Template) Parse(args ...interface{}) string {
	var buf bytes.Buffer
	buf.Grow(t.size)

	t.ParseOutput(&buf, args...)
	return buf.String()
}

func Parse(text string, args ...interface{}) string {
	return NewTemplate(text).Parse(args...)
}

func (t *Template) ParseOutput(w io.Writer, args ...interface{}) (err error) {
	t.Reset()
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// wrap w.Write func, handle error
	write := func(w io.Writer, p []byte) {
		if _, err := w.Write(p); err != nil {
			panic(err)
		}
	}

	var index int
	for {
		// stop until meet '{'
		pos := t.pos
		for pos < t.size && t.chars[pos] != t.begin {
			pos++
		}
		if pos <= t.size {
			write(w, t.chars[t.pos:pos])
			t.pos = pos
		}

		// read next char
		ch, ok := t.nextChar(false)
		if !ok {
			break
		}

		// handle char '{'
		if ch == t.begin && index < len(args) {
			next, ok := t.nextChar(true)
			if !ok {
				write(w, []byte{ch})
				break
			}

			// handle char '}'
			if next == t.end {
				t.nextChar(false) // eat '}'
				write(w, []byte(ConvertToString(args[index])))
				index++
			} else {
				write(w, []byte{ch})
			}
		} else {
			write(w, []byte{ch})
		}
	}
	return nil
}

func ParseOutput(w io.Writer, text string, args ...interface{}) error {
	return NewTemplate(text).ParseOutput(w, args...)
}

func ConvertToString(v interface{}) string {
	if v == nil {
		return "nil"
	}

	switch v := v.(type) {
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case uintptr:
		return "0x" + strconv.FormatUint(uint64(v), 16)
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case []byte:
		return string(v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%+v", v)
	}
}
