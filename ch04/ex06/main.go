package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	bytes := []byte("         Hello     ,  \t \t \t World!")
	fmt.Println(string(compress(bytes)))
}

func compress(bytes []byte) []byte {
	i := 0
	for bp := 0; bp < len(bytes); {
		b, size := utf8.DecodeRune(bytes[bp:])
		if !unicode.IsSpace(b) {
			utf8.EncodeRune(bytes[i:], b)
			i += size
		}
		bp += size
	}
	return bytes[:i]
}
