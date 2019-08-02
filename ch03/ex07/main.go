package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	iter     = 10
	contrast = 2
	maxcmpl  = 1e-5
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2.0, -2.0, +2.0, +2.0
		width, height          = 256, 256
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
	for i := uint8(0); i < iter; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < maxcmpl {

			r, g, b, a := archtan(z).RGBA()
			return color.RGBA{uint8(r) / (i + 1), uint8(g) / (i + 1), uint8(b) / (i + 1), uint8(a)}
		}
	}
	return color.Black
}
func archtan(z complex128) color.Color {
	return color.YCbCr{192, uint8(real(cmplx.Atan(z))), uint8(imag(cmplx.Atan(z)))}
}
