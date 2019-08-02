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
			input:  []byte("hello,           world !!   "),
			expect: []byte("hello,world!!"),
		},
		{
			input:  []byte("hello,     \t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t      world !!   "),
			expect: []byte("hello,world!!"),
		},
		{
			input:  []byte("hello,   \t\t\t \t\t\t\t        world !!   "),
			expect: []byte("hello,world!!"),
		},
	}
	for i, test := range tests {
		result := compress(test.input)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("No: %d\nresult:\n%v\nexpect:\n%v\n", i+1, result, test.expect)
		}
	}
}
