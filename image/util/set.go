package util

import (
	"image"
	"image/color"
)

type Setable interface {
	image.Image
	Set(x, y int, c color.Color)
}

func Copy(in image.Image, out Setable) {
	minInX := in.Bounds().Min.X
	minInY := in.Bounds().Min.Y
	minOutX := out.Bounds().Min.X
	minOutY := out.Bounds().Min.Y
	dimX := out.Bounds().Max.X - minOutX
	dimY := out.Bounds().Max.Y - minOutY

	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			out.Set(x+minOutX, y+minOutY, in.At(minInX+x, minInY+y))
		}
	}
}
