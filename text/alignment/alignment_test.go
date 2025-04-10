package alignment

import (
	"fmt"
	"testing"
)

var (
	s = `host :=        localhost
  p := 3306
   init-sql         := "set names utf8"
`
)

func TestAlignment(t *testing.T) {
	var testcases map[string]string = map[string]string{
		":=": `host     := localhost
p        := 3306
init-sql := "set names utf8"
`,
		"=": `host :             = localhost
p :                = 3306
init-sql         : = "set names utf8"
`,
	}
	for sep, expected := range testcases {
		if res := NewAlignment(sep).Format(s); res != expected {
			t.Fatalf("sep = %s, the result did not meet expectations", sep)
		} else {
			fmt.Println(res)
		}
	}
}
