package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	rs := []rune{' ', '\t', 'H', 'e', 'l', 'l', 'o', '\t', ' ', '　', '\t', 'w', 'o', 'r', 'l', 'd', '\r', '\v'}
	bytes := []byte(string(rs))
	fmt.Println(string(compressSpace(bytes)))
}

func compressSpace(bytes []byte) []byte {
	i := 0
	for j := 0; j < len(bytes); {
		r, size := utf8.DecodeRune(bytes[j:])
		last, _ := utf8.DecodeLastRune(bytes[:i])
		if !unicode.IsSpace(last) || !unicode.IsSpace(r) {
			if unicode.IsSpace(r) {
				r = ' '
			}
			_ = utf8.EncodeRune(bytes[i:], r)
			i += size
		}
		j += size
	}
	return bytes[:i]
}
