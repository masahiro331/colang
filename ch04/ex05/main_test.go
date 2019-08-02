package main

import (
	"reflect"
	"testing"
)

func TestUniq(t *testing.T) {
	var tests = []struct {
		input  []string
		expect []string
	}{
		{
			input:  []string{"apple", "beer", "beer", "apple"},
			expect: []string{"apple", "beer", "apple"},
		},
		{
			input:  []string{"ab", "ab", "ba", "ba", "cd"},
			expect: []string{"ab", "ba", "cd"},
		},
		{
			input:  []string{"ab", "ab", "ba", "ba", "cd", "cd", "cd"},
			expect: []string{"ab", "ba", "cd"},
		},
		{
			input:  []string{"a", "b", "c"},
			expect: []string{"a", "b", "c"},
		},
	}
	for i, test := range tests {
		result := uniq(test.input)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("No: %d\nresult:\n%v\nexpect:\n%v\n", i+1, result, test.expect)
		}
	}
}
