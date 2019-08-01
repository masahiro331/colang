package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("100"))
}

func comma(s string) string {
	sl := len(s)
	if sl < 4 {
		return s
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
		return ret[1:]
	}
	return ret
}
