package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{name: "empty body", input: "<html><head></head><body></body></html>", expected: map[string]int{"html": 1, "head": 1, "body": 1}},
		{name: "some tags",
			input: `
<html>
	<head>
	</head>
	<body>
		<a href="example.com">aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</a>
		<div class="fuga">
			<p>fuga</p>
		</div>
	</body>
</html>`,
			expected: map[string]int{"html": 1, "head": 1, "body": 1, "a": 1, "div": 1, "p": 1}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			actual := make(map[string]int)
			outline(actual, doc)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("input %s expect %v actual %v", testCase.input, testCase.expected, actual)
			}

		})

	}
}
