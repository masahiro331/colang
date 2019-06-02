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
			"hello\n",
		},
		{
			[]string{
				"./echo",
				"hello",
				"world",
			},
			"hello world\n",
		},
		{
			[]string{
				"./echo",
				"Hello",
				"World",
			},
			"Hello World\n",
		},
		{
			[]string{
				"./echo",
				"Hello",
				"World",
				"!!!",
			},
			"Hello World !!!\n",
		},
	}
	for i, test := range tests {
		out = new(bytes.Buffer)
		if err := echoWithJoin(test.args); err != nil {
			t.Errorf("echo %+v failed: %v", test.args, err)
			continue
		}
		result := out.(*bytes.Buffer).String()
		if result != test.expectOut {
			t.Errorf("No: %d\nresult:\n%s\nexpect:\n%s\n", i+1, result, test.expectOut)
		}
	}
}

func BenchmarkEcho(b *testing.B) {
	testArgs := []string{"Hello", "World", "!!!"}
	out = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		echo(testArgs)
	}
}

func BenchmarkEchoWithJoin(b *testing.B) {
	testArgs := []string{"Hello", "World", "!!!"}
	out = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		echoWithJoin(testArgs)
	}
}
