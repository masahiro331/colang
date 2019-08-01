package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		input  string
		expect string
	}{
		{
			input:  "100",
			expect: "100",
		},
		{
			input:  "1000",
			expect: "1,000",
		},
		{
			input:  "10000",
			expect: "10,000",
		},
		{
			input:  "100000",
			expect: "100,000",
		},
		{
			input:  "1000000",
			expect: "1,000,000",
		},
		{
			input:  "1000000.44444",
			expect: "1,000,000.44444",
		},
	}
	for i, test := range tests {
		if result := comma(test.input); result != test.expect {
			t.Errorf("No: %d\nresult:\n%s\nexpect:\n%s\n", i+1, result, test.expect)
		}
	}
}
