package main

import (
	"image"
)

func sameSize(in image.Image) *image.RGBA {
	bounds := in.Bounds()
	outbounds := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: bounds.Max.X - bounds.Min.X,
			Y: bounds.Max.Y - bounds.Min.Y,
		},
	}
	out := image.NewRGBA(outbounds)
	return out
}

func copy(in image.Image) *image.RGBA {
	out := sameSize(in)
	minX := in.Bounds().Min.X
	minY := in.Bounds().Min.Y
	for x := 0; x < out.Bounds().Max.X; x++ {
		for y := 0; y < out.Bounds().Max.Y; y++ {
			out.Set(x, y, in.At(minX+x, minY+y))
		}
	}
	return out
}
