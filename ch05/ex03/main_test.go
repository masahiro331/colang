package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestFindtexts(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "only title", input: "<html><head><title>TITLE</title></head><body></body></html>", expected: []string{"TITLE"}},
		{name: "only body", input: "<html><head></head><body><h1>BODY</h1></body></html>", expected: []string{"BODY"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			actual := findtexts(nil, doc)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("input %s expect %v actual %v", testCase.input, testCase.expected, actual)
			}

		})

	}
}
