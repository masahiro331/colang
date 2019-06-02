package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	if err := echo(os.Args); err != nil {
		log.Fatal(err)
	}
}

func echo(args []string) error {
	s, sep := "", ""
	for i, arg := range args[1:] {
		s += fmt.Sprintf("%s%d %s", sep, i+1, arg)
		sep = "\n"
	}
	_, err := fmt.Fprintln(out, s)
	if err != nil {
		return err
	}
	return nil
}
