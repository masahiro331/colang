package main

import (
	"github.com/masahiro331/colang/ch12/ex02/display"
)

type Hoge struct {
	Name string
	Hoge *Hoge
}

func main() {
	hoge := Hoge{
		Name: "hoge",
		Hoge: &Hoge{
			Name: "fuga",
			Hoge: &Hoge{
				Name: "fuga",
				Hoge: &Hoge{
					Name: "fuga",
					Hoge: &Hoge{
						Name: "fuga",
						Hoge: &Hoge{
							Name: "fuga",
							Hoge: &Hoge{
								Name: "fuga",
							},
						},
					},
				},
			},
		},
	}
	display.Display("struct", hoge)
}
