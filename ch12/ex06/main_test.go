package main

import (
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	var tests = []struct {
		name  string
		input interface{}
		want  string
	}{
		{name: "bool true", input: true, want: "t"},
		{name: "bool false", input: false, want: "nil"},
		{name: "float positive zero", input: 0.0, want: "0.000000"},
		{name: "float negative zero", input: -0.0, want: "0.000000"},
		{name: "interface nil", input: struct{ v interface{} }{nil}, want: `((v nil))`},
		{name: "interface array", input: struct{ v interface{} }{[]int{1, 2, 3}}, want: `((v ("[]int" (1 2 3))))`},
		{name: "zero check 1 value", input: struct{ x int }{0}, want: "()"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Marshal(test.input)
			if err != nil || string(got) != test.want {
				t.Errorf("Marshal(%v) = \n%q, %v\n want %q, nil", test.input, got, err, test.want)
			}
		})
	}
}
func TestIsZeroValue(t *testing.T) {
	var zeroTests = []struct {
		input interface{}
		want  bool
	}{
		{false, true},
		{true, false},
		{int(0), true},
		{int(1), false},
		{uint(0), true},
		{uint(1), false},
		{"", true},
		{"a", false},
	}
	for _, test := range zeroTests {
		got := isZeroValue(reflect.ValueOf(test.input))
		if got != test.want {
			t.Errorf("isZeroValue(%v) = %t\n want %t", test.input, got, test.want)
		}
	}
}
