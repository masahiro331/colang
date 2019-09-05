package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "empty body", input: "<html><head></head><body></body></html>", expected: nil},
		{name: "2 link body",
			input: `
<html>
	<head>
	</head>
	<body>
		<a href="example..com">aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</a>
		<a href="/top">top</a>
	</body>
</html>`,
			expected: []string{"example.com", "/top"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			actual := visit(nil, doc)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("input %s expect %v actual %v", testCase.input, testCase.expected, actual)
			}

		})

	}
}
