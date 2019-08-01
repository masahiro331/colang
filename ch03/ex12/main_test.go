package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		expect bool
	}{
		{
			input1: "Hello, World",
			input2: "Hlo,elrl Wod",
			expect: true,
		},
		{
			input1: "Hello, 世界",
			input2: "He世l, 界lo",
			expect: true,
		},
		{
			input1: "Hello, World",
			input2: "Hlo,lrl Wod",
			expect: false,
		},
		{
			input1: "Hello, 世界",
			input2: "He世l 界lo",
			expect: false,
		},
	}
	for i, test := range tests {
		if result := isAnagram(test.input1, test.input2); result != test.expect {
			t.Errorf("No: %d\nresult:\n%v\nexpect:\n%v\n", i+1, result, test.expect)
		}
	}
}
