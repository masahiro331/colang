package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		vsize                  = 2
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		vwidth, vheight        = width * vsize, height * vsize
	)

	vimage := make([][]uint8, vheight)
	for i := 0; i < vheight; i++ {
		vimage[i] = make([]uint8, vwidth)
	}

	for py := 0; py < vheight; py++ {
		y := float64(py)/vheight*(ymax-ymin) + ymin
		for px := 0; px < vwidth; px++ {
			x := float64(px)/vwidth*(xmax-xmin) + xmin
			z := complex(x, y)
			vimage[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			sum := vimage[px*2][py*2] +
				vimage[px*2+0][py*2+1] +
				vimage[px*2+1][py*2+0] +
				vimage[px*2+1][py*2+1]

			img.Set(px, py, color.Gray{uint8(sum / 4)})
		}
	}

	png.Encode(os.Stdout, img)
}

// GrayImage
func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return uint8(255 - contrast*n)
		}
	}

	return uint8(0)
}
