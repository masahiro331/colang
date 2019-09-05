package main

import "fmt"

func main() {
	fmt.Println(echo(10))
}

func echo(i int) (res int) {
	defer func() {
		recover()
		res = 1
	}()
	panic(nil)
}
