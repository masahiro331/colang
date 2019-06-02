package main

import (
	"bytes"
	"image"
	"testing"
)

func TestLissajous(t *testing.T) {
	out := new(bytes.Buffer)
	lissajous(out)

	_, _, err := image.DecodeConfig(out)
	if err != nil {
		t.Errorf("failed to decode out file: %v", err)
	}
}
