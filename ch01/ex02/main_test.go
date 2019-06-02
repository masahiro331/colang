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
			"1 hello\n",
		},
		{
			[]string{
				"./echo",
				"hello",
				"world",
			},
			"1 hello\n2 world\n",
		},
		{
			[]string{
				"./echo",
				"Hello",
				"World",
			},
			"1 Hello\n2 World\n",
		},
		{
			[]string{
				"./echo",
				"Hello",
				"World",
				"!!!",
			},
			"1 Hello\n2 World\n3 !!!\n",
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
			t.Errorf("No: %d\nresult:\n%s\nexpect:\n%s\n", i+1, result, test.expectOut)
		}
	}
}
