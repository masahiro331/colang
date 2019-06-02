package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	var tests = []struct {
		url          string
		expectHeader string
	}{
		{
			url:          "http://example.com",
			expectHeader: "200 OK",
		},
		{
			url:          "http://example.com/notfound",
			expectHeader: "404 Not Found",
		},
	}

	for i, test := range tests {
		out = new(bytes.Buffer)
		err := fetch(test.url)
		if err != nil {
			t.Errorf("Error: %s", test.url)
		}

		result := out.(*bytes.Buffer).String()
		scanner := bufio.NewScanner(strings.NewReader(result))
		if scanner.Scan() {
			status := scanner.Text()
			if status != test.expectHeader {
				t.Errorf("No: %d,want: %s, actual: %s", i, test.expectHeader, status)
			}
		}
	}
}
