package main

import (
	"regexp"
	"testing"
)

func TestFetch(t *testing.T) {
	var tests = []struct {
		url               string
		ch                chan string
		expectResultRegex string
	}{
		{
			url:               "http://example.com",
			ch:                make(chan string),
			expectResultRegex: `[0-9]\.[0-9]{2}s [0-9|\ ]{7} http://example.com`,
		},
	}

	for i, test := range tests {
		var result string
		r := regexp.MustCompile(test.expectResultRegex)
		go fetch(test.url, test.ch)
		result = <-test.ch
		if !r.MatchString(result) {
			t.Errorf("No: %d,want: %s, actual: %s", i, test.expectResultRegex, result)
		}
	}
}
