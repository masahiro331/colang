package main

import (
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	var tests = []struct {
		input  []byte
		expect []byte
	}{
		{
			input:  []byte("hello,world!!"),
			expect: []byte("!!dlrow,olleh"),
		},
		{
			input:  []byte("Go 言語　実習"),
			expect: []byte("習実　語言 oG"),
		},
	}
	for i, test := range tests {
		Reverse(test.input)
		if !reflect.DeepEqual(test.input, test.expect) {
			t.Errorf("No: %d\nresult:\n%v\nexpect:\n%v\n", i+1, test.input, test.expect)
		}
	}
}
