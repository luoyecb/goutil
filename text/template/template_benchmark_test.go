package template

import (
	"net/url"
	"testing"

	"github.com/valyala/fasttemplate"
)

func BenchmarkMyTemplate(b *testing.B) {
	t := NewTemplate(tpl)
	args := []interface{}{"www.glc.com", TplString("hello"), "foobar", TplStruct{0, "hello"}}
	for i := 0; i < b.N; i++ {
		t.Parse(args...)
	}
}

func BenchmarkFaskTemplate(b *testing.B) {
	tpl := "http://{{host}}/?q={{query}}&foo={{bar}}&unknow_key={{unknow}}"
	placeholders := map[string]interface{}{
		"host":  "www.glc.com",
		"query": url.QueryEscape("hello=world"),
		"bar":   "foobar",
	}

	t := fasttemplate.New(tpl, "{{", "}}")
	for i := 0; i < b.N; i++ {
		t.ExecuteString(placeholders)
	}
}
