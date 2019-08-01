package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	iterations = 30
	contrast   = 5
	maxcmpl    = 1e-6
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2.0, -2.0, +2.0, +2.0
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < maxcmpl {
			c := archcos(z)
			r, g, b, a := c.RGBA()
			return color.RGBA{uint8(r) / (i + 1), uint8(g) / (i + 1), uint8(b) / (i + 1), uint8(a)}
		}
	}
	return color.Black
}
func archcos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}
