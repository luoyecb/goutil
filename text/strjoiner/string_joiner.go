package strjoiner

import (
	"bytes"
)

type StringJoiner struct {
	Delimiter  string // 分隔符
	Prefix     string
	Suffix     string
	PrefixElem string
	SuffixElem string

	value  bytes.Buffer
	length int
}

func NewStringJoiner(d string) *StringJoiner {
	return &StringJoiner{Delimiter: d}
}

func NewStringJoiner3(d, prefix, suffix string) *StringJoiner {
	return &StringJoiner{
		Delimiter: d,
		Prefix:    prefix,
		Suffix:    suffix,
	}
}

func NewStringJoiner5(d, prefix, suffix, prefixElem, suffixElem string) *StringJoiner {
	return &StringJoiner{
		Delimiter:  d,
		Prefix:     prefix,
		Suffix:     suffix,
		PrefixElem: prefixElem,
		SuffixElem: suffixElem,
	}
}

func (sj *StringJoiner) Len() int {
	return sj.length
}

func (sj *StringJoiner) updateLen() {
	l := sj.value.Len()
	if l > 0 {
		l += len(sj.Prefix) + len(sj.Suffix)
	}
	sj.length = l
}

func (sj *StringJoiner) Add(s string) *StringJoiner {
	sj.add(func() {
		sj.value.WriteString(s)
	})
	sj.updateLen()
	return sj
}

func (sj *StringJoiner) AddBytes(bytes []byte) *StringJoiner {
	sj.add(func() {
		sj.value.Write(bytes)
	})
	sj.updateLen()
	return sj
}

func (sj *StringJoiner) add(fn func()) {
	if sj.Len() > 0 {
		sj.value.WriteString(sj.Delimiter)
	}

	if sj.PrefixElem != "" {
		sj.value.WriteString(sj.PrefixElem)
	}
	fn()
	if sj.SuffixElem != "" {
		sj.value.WriteString(sj.SuffixElem)
	}
}

func (sj *StringJoiner) String() string {
	if sj.Len() > 0 {
		return sj.Prefix + sj.value.String() + sj.Suffix
	}
	return ""
}
