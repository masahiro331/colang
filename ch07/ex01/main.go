package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	var c WordCounter
	fmt.Fprintf(&c, "hogehoge fuga fuga")
	fmt.Println(c)

	var c2 LineCounter
	fmt.Fprintf(&c2, "hogehoge fuga fuga")
	fmt.Println(c2)
}
