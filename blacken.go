package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"
)

type blacken struct {
	minPercentile      int
	maxPercentile      int
	blackenProbability float64
}

// given an image, blacken the pixels between the specified percentiles with
// the specified probability, generating a new image.
func (b *blacken) transform(in image.Image) image.Image {

	brightnessSort := newBrightnessSort(in)
	minX := b.minPercentile * len(brightnessSort.permutation) / 100
	maxX := b.maxPercentile * len(brightnessSort.permutation) / 100

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := copy(in)

	for x := minX; x < maxX; x++ {
		p := brightnessSort.permutation[x]
		if r.Float64() < b.blackenProbability {
			out.Set(p.X, p.Y, color.Black)
		}
	}
	return out
}
