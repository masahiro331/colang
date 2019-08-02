package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bytes := []byte("Hello, World!")
	Reverse(bytes)
	fmt.Println(string(bytes))
}

func reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func Reverse(bytes []byte) {
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		reverse(bytes[i : i+size])
		i += size
	}
	reverse(bytes)
}
