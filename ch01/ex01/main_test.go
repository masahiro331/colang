package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args      []string
		expectOut string
	}{
		{
			[]string{
				"./echo",
				"hello",
			},
			"./echo hello\n",
		},
		{
			[]string{
				"./echo",
				"hello",
				"world",
			},
			"./echo hello world\n",
		},
		{
			[]string{
				"./echo",
				"Hello",
				"World",
			},
			"./echo Hello World\n",
		},
		{
			[]string{
				"./echo",
				"Hello",
				"World",
				"!!!",
			},
			"./echo Hello World !!!\n",
		},
	}
	for i, test := range tests {
		out = new(bytes.Buffer)
		if err := echo(test.args); err != nil {
			t.Errorf("echo %+v failed: %v", test.args, err)
			continue
		}
		result := out.(*bytes.Buffer).String()
		if result != test.expectOut {
			t.Errorf("No: %d, result: %s, expect: %s", i+1, result, test.expectOut)
		}
	}
}
