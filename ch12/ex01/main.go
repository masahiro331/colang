package main

import (
	"github.com/masahiro331/colang/ch12/ex01/display"
)

func main() {
	display.Display("struct", struct {
		Name string
		Sex  string
	}{
		Name: "hoge",
		Sex:  "man",
	})
}
