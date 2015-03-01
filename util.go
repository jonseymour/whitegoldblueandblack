package main

import (
	"image"
)

func sameSize(in image.Image) *image.RGBA64 {
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
	out := image.NewRGBA64(outbounds)
	return out
}
