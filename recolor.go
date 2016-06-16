package main

import (
	"github.com/jonseymour/whitegoldblueandblack/image/util"
	"image"
	"image/color"
)

func compare(sample color.Color, ref color.Color) bool {
	lR, lG, lB, _ := sample.RGBA()
	rR, rG, rB, _ := ref.RGBA()
	return lR == rR && lG == rG && lB == rB
}

func recolor(in image.Image, ref color.Color, replace color.Color) image.Image {

	out := sameSize(in)
	util.Copy(in, out)

	for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
		for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
			if compare(in.At(x, y), ref) {
				out.Set(x, y, replace)
			}
		}
	}
	return out
}
