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
			url:         "https://example.com",
			expectError: "fetch: Get http://https//example.com: dial tcp: lookup https: no such host\n",
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

func TestResolveUrl(t *testing.T) {
	var tests = []struct {
		inputUrl  string
		expectUrl string
	}{
		{
			inputUrl:  "example.com",
			expectUrl: "http://example.com",
		},
		{
			inputUrl:  "https://example.com",
			expectUrl: "http://https://example.com",
		},
		{
			inputUrl:  "http://example.com",
			expectUrl: "http://example.com",
		},
	}
	for i, test := range tests {
		if resolveUrl(test.inputUrl) != test.expectUrl {
			t.Errorf("No: %d,want: %s, actual: %s", i, test.expectUrl, resolveUrl(test.inputUrl))
		}
	}

}
