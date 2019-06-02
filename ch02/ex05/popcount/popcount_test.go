package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	var testCases = []struct {
		input  uint64
		expect int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			7,
			3,
		},
		{
			15,
			4,
		},
		{
			31,
			5,
		},
		{
			63,
			6,
		},
		{
			127,
			7,
		},
		{
			255,
			8,
		},
		{
			511,
			9,
		},
		{
			1023,
			10,
		},
		{
			1024,
			1,
		},
	}
	for i, test := range testCases {
		if PopCount(test.input) != test.expect {
			t.Errorf("No: %d, want: %d, expect %d", i+1, PopCount(test.input), test.expect)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1023)
	}
}

func BenchmarkPopCountOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOld(1023)
	}
}
