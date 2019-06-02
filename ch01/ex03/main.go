package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	if err := echoWithJoin(os.Args); err != nil {
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

func echoWithJoin(args []string) error {
	s, sep := "", " "
	s = strings.Join(args[1:], sep)

	_, err := fmt.Fprintln(out, s)
	if err != nil {
		return err
	}
	return nil
}
