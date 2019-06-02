package main

import (
	"reflect"
	"testing"
)

func TestDup(t *testing.T) {
	var tests = []struct {
		file      string
		expectMap map[string]int
	}{
		{
			file: "test_cases/test1.txt",
			expectMap: map[string]int{
				"hoge": 2,
				"piyo": 2,
				"fuga": 2,
			},
		},
		{
			file: "test_cases/test2.txt",
			expectMap: map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
				"4": 4,
			},
		},
		{
			file:      "test_cases/test3.txt",
			expectMap: map[string]int{},
		},
	}
	for i, test := range tests {
		result, err := dup(test.file)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if !reflect.DeepEqual(result, test.expectMap) {
			t.Errorf("No: %d, want: %v, actual: %v", i+1, result, test.expectMap)
		}
	}
}
