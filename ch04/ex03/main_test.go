package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		input  [5]int
		expect [5]int
	}{
		{
			input:  [...]int{1, 2, 3, 4, 5},
			expect: [...]int{5, 4, 3, 2, 1},
		},
	}
	for i, test := range tests {
		reverse(&test.input)
		if test.input != test.expect {
			t.Errorf("No: %d\nresult:\n%v\nexpect:\n%v\n", i+1, test.input, test.expect)
		}
	}
}
