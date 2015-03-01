package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"
)

type colorize struct {
	minPercentile       int
	maxPercentile       int
	colorizeProbability float64
	color               color.Color
	refColor            color.Color
}

// given an image, colorize the pixels between the specified percentiles with
// the specified probability, generating a new image.
func (b *colorize) transform(in image.Image) image.Image {

	brightnessSort := newBrightnessSort(in, b.refColor)
	minX := b.minPercentile * len(brightnessSort.permutation) / 100
	maxX := b.maxPercentile * len(brightnessSort.permutation) / 100

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := copy(in)

	for x := minX; x < maxX; x++ {
		p := brightnessSort.permutation[x]
		if r.Float64() < b.colorizeProbability {
			out.Set(p.X, p.Y, b.color)
		}
	}
	return out
}
