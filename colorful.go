package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
)

// given an image in an arbitrary color space, transform it into the CIELab color space, using colorful
func toColorful(in color.Color) colorful.Color {
	r, g, b, _ := in.RGBA()
	x, y, z := colorful.LinearRgbToXyz(float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0)
	return colorful.Xyz(x, y, z)
}
