package main

import (
	"bytes"
	"image"
	"testing"
)

func TestLissajous(t *testing.T) {
	cycles := 10
	out := new(bytes.Buffer)
	lissajous(out, cycles)

	_, _, err := image.DecodeConfig(out)
	if err != nil {
		t.Errorf("failed to decode out file: %v", err)
	}
}
