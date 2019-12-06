package sexpr

import "testing"

func TestMarshal(t *testing.T) {
	var tests = []struct {
		name  string
		input interface{}
		want  string
	}{
		{name: "bool true", input: true, want: "t"},
		{name: "bool false", input: false, want: "nil"},
		{name: "complex both positive", input: 1 + 2i, want: "#C(1 2)"},
		{name: "complex real positive imag negative", input: 2 - 1i, want: "#C(2 -1)"},
		{name: "complex real negative imag positive", input: -3 + 4i, want: "#C(-3 4)"},
		{name: "complex real negative imag negative", input: -1 - 2i, want: "#C(-1 -2)"},
		{name: "interface nil", input: struct{ v interface{} }{nil}, want: `((v nil))`},
		{name: "interface array", input: struct{ v interface{} }{[]int{1, 2, 3}}, want: `((v ("[]int" (1 2 3))))`},
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
