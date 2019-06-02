package main

import (
	"testing"
)

func TestFetch(t *testing.T) {
	var error_tests = []struct {
		url         string
		expectError string
	}{
		{
			url:         "example.com",
			expectError: "fetch: Get example.com: unsupported protocol scheme \"\"\n",
		},
	}
	for i, test := range error_tests {
		err := fetch(test.url)
		if err == nil {
			t.Errorf("Error: %s", test.url)
		}
		if err != nil && err.Error() != test.expectError {
			t.Errorf("No: %d,want: %s, actual: %s", i, test.expectError, err.Error())
		}
	}
}
