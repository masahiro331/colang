package tempconv

import (
	"testing"
)

type testCase struct {
	description string
	input       float64
	expect      float64
}

func TestTempconv(t *testing.T) {
	var testsKToF = []testCase{
		{
			"Test KToF",
			100,
			-279.67,
		},
	}
	var testsKToC = []testCase{
		{
			"Test KToC",
			100,
			-173.14999999999998,
		},
	}
	var testsFToC = []testCase{
		{
			"Test FToC",
			140,
			60,
		},
	}
	var testsFToK = []testCase{
		{
			"Test FToK",
			100,
			310.92777777777775,
		},
	}
	var testsCToF = []testCase{
		{
			"Test CToF",
			100,
			212,
		},
	}
	var testsCToK = []testCase{
		{
			"Test CToK",
			0,
			273.15,
		},
	}

	for i, test := range testsKToF {
		if KToF(Kelvin(test.input)) != Fahrenheit(test.expect) {
			t.Errorf("No: %d Test Description: %s \nwant: %f, expect: %f", i+1, test.description, KToF(Kelvin(test.input)), test.expect)
		}
	}
	for i, test := range testsKToC {
		if KToC(Kelvin(test.input)) != Celsius(test.expect) {
			t.Errorf("No: %d Test Description: %s \nwant: %f, expect: %f", i+1, test.description, KToC(Kelvin(test.input)), test.expect)
		}
	}
	for i, test := range testsFToC {
		if FToC(Fahrenheit(test.input)) != Celsius(test.expect) {
			t.Errorf("No: %d Test Description: %s \nwant: %f, expect: %f", i+1, test.description, FToC(Fahrenheit(test.input)), test.expect)
		}
	}
	for i, test := range testsFToK {
		if FToK(Fahrenheit(test.input)) != Kelvin(test.expect) {
			t.Errorf("No: %d Test Description: %s \nwant: %f, expect: %f", i+1, test.description, FToK(Fahrenheit(test.input)), test.expect)
		}
	}
	for i, test := range testsCToK {
		if CToK(Celsius(test.input)) != Kelvin(test.expect) {
			t.Errorf("No: %d Test Description: %s \nwant: %f, expect: %f", i+1, test.description, CToK(Celsius(test.input)), test.expect)
		}
	}
	for i, test := range testsCToF {
		if CToF(Celsius(test.input)) != Fahrenheit(test.expect) {
			t.Errorf("No: %d Test Description: %s \nwant: %f, expect: %f", i+1, test.description, CToF(Celsius(test.input)), test.expect)
		}
	}
}
