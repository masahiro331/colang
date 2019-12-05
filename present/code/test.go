package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var out io.Writer = os.Stdout

func main() {
	Echo()
}

func Echo() {
	_, _ := fmt.Fprintln(out, "HOGEHOGE")
	return nil
}

func TestEcho(t *testing.T) {
	out = new(bytes.Buffer)
	Echo()
	result := out.(*bytes.Buffer).String()
	if out != "HOGEHOGE" {
		t.Errorf("errorj")
	}
}
