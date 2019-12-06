package main

import (
	"github.com/masahiro331/colang/ch12/ex01/display"
)

type Hoge struct {
	Name string
}

func main() {
	display.Display("struct", struct {
		Name  string
		Sex   string
		Array []string
		Hoges []Hoge
	}{
		Name: "hoge",
		Sex:  "man",
		Array: []string{
			"hoge",
			"hoge",
		},
		Hoges: []Hoge{
			{
				Name: "hoge",
			},
		},
	})
}
