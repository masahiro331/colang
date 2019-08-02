package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	rs := []rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', 'H', 'e', 'l', 'l', 'o', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', '\t', ' ', '　', '\t', 'w', 'o', 'r', 'l', 'd', '!'}
	bytes := []byte(string(rs))
	fmt.Println(string(cs(bytes)))
}

func cs(bytes []byte) []byte {
	i := 0
	for j := 0; j < len(bytes); {
		b, size := utf8.DecodeRune(bytes[j:])

		last, _ := utf8.DecodeLastRune(bytes[:i])
		if !unicode.IsSpace(last) || !unicode.IsSpace(b) {
			utf8.EncodeRune(bytes[i:], b)
			i += size
		}
		j += size
	}
	return bytes[:i]
}
