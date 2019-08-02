package main

import (
	"reflect"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		input  []int
		rotate int
		expect []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			rotate: 3,
			expect: []int{3, 4, 5, 1, 2},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			rotate: 4,
			expect: []int{2, 3, 4, 5, 1},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			rotate: 5,
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			rotate: 6,
			expect: []int{5, 1, 2, 3, 4},
		},
	}
	for i, test := range tests {
		result := rotate(test.input, test.rotate)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("No: %d\nresult:\n%v\nexpect:\n%v\n", i+1, result, test.expect)
		}
	}
}
