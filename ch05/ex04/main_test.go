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
		{name: "simple html",
			input: `
<html>
	<head>
		<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
	</head>
	<body>
		<a href="example.com">aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</a>
		<a href="/top">top</a>
		<img class="img" src="sample.gif">
	</body>
</html>`,
			expected: []string{"/lib/godoc/style.css", "https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js", "example.com", "/top", "sample.gif"}},
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
