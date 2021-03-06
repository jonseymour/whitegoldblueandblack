package main

import (
	distance "github.com/jonseymour/whitegoldblueandblack/image/color"
	"github.com/jonseymour/whitegoldblueandblack/image/util"
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
	metric              distance.DistanceMetric
}

// given an image, colorize the pixels between the specified percentiles with
// the specified probability, generating a new image.
func (b *colorize) transform(in image.Image) image.Image {

	distanceSort := newDistanceSort(in, b.refColor, b.metric)
	minX := b.minPercentile * len(distanceSort.permutation) / 100
	maxX := b.maxPercentile * len(distanceSort.permutation) / 100

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := sameSize(in)
	util.Copy(in, out)

	for x := minX; x < maxX; x++ {
		p := distanceSort.permutation[x]
		if r.Float64() < b.colorizeProbability {
			out.Set(p.X, p.Y, b.color)
		}
	}
	return out
}
