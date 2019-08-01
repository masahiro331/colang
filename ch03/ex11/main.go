package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("100"))
}

func comma(s string) string {
	var as string
	sl := len(s)
	if sl < 4 {
		return s
	}

	if i := strings.Index(s, "."); i > 0 {
		as = s[i:]
		s = s[:i]
	}

	var buf bytes.Buffer
	buf.WriteString(s[:sl%3])
	for i, c := range s[sl%3:] {
		if i%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}

	ret := buf.String()
	if string(ret[0]) == "," {
		return ret[1:] + as
	}
	return ret + as
}
