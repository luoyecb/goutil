package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tpl  = "http://{}?q={}&foo={}&unknow_key={}"
	tpl2 = "http://{}?q={{}"
	tpl3 = "http://www.glc.com"
)

type TplString string

type TplStruct struct {
	id int
	s  string
}

type testCase struct {
	expected string
	args     []interface{}
}

func TestParse1(t *testing.T) {
	assert := assert.New(t)

	tt := NewTemplate(tpl)
	testCases := []*testCase{
		&testCase{"http://{}?q={}&foo={}&unknow_key={}", []interface{}{}},
		&testCase{"http://www.glc.com?q={}&foo={}&unknow_key={}", []interface{}{"www.glc.com"}},
		&testCase{"http://www.glc.com?q=hello&foo={}&unknow_key={}", []interface{}{"www.glc.com", TplString("hello")}},
		&testCase{"http://www.glc.com?q=hello&foo=foobar&unknow_key={}", []interface{}{"www.glc.com", TplString("hello"), "foobar"}},
		&testCase{"http://www.glc.com?q=hello&foo=foobar&unknow_key=nil", []interface{}{"www.glc.com", TplString("hello"), "foobar", nil}},
		&testCase{"http://www.glc.com?q=hello&foo=foobar&unknow_key={id:0 s:hello}", []interface{}{"www.glc.com", TplString("hello"), "foobar", TplStruct{0, "hello"}}},
		&testCase{"http://www.glc.com?q=hello&foo=foobar&unknow_key=&{id:0 s:hello}", []interface{}{"www.glc.com", TplString("hello"), "foobar", &TplStruct{0, "hello"}}},
		// &testCase{[]interface{}{}},
	}
	for _, arg := range testCases {
		assert.Equal(arg.expected, tt.Parse(arg.args...))
	}
}

func TestParse2(t *testing.T) {
	assert := assert.New(t)

	tt := NewTemplate(tpl2)
	testCases := []*testCase{
		&testCase{"http://{}?q={{}", []interface{}{}},
		&testCase{"http://www.glc.com?q={{}", []interface{}{"www.glc.com"}},
		&testCase{"http://www.glc.com?q={hello", []interface{}{"www.glc.com", TplString("hello")}},
		&testCase{"http://www.glc.com?q={hello", []interface{}{"www.glc.com", TplString("hello"), "foobar"}},
	}
	for _, arg := range testCases {
		assert.Equal(arg.expected, tt.Parse(arg.args...))
	}
}

func TestParse3(t *testing.T) {
	assert := assert.New(t)

	tt := NewTemplate(tpl3)
	testCases := []*testCase{
		&testCase{"http://www.glc.com", []interface{}{}},
		&testCase{"http://www.glc.com", []interface{}{"www.glc.com"}},
		&testCase{"http://www.glc.com", []interface{}{"www.glc.com", TplString("hello")}},
		&testCase{"http://www.glc.com", []interface{}{"www.glc.com", TplString("hello"), "foobar"}},
	}
	for _, arg := range testCases {
		assert.Equal(arg.expected, tt.Parse(arg.args...))
	}
}
