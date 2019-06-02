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
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	_, err := fmt.Fprintln(out, s)
	if err != nil {
		return err
	}
	return nil
}
