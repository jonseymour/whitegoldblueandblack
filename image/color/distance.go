package color

import (
	"image/color"
	"math"
)

// A function that can calculate a distane metric between two colors.
type DistanceMetric func(from color.Color, to color.Color) float64

func RGBADistanceMetric(from color.Color, to color.Color) float64 {
	r, g, b, _ := from.RGBA()
	rr, rg, rb, _ := to.RGBA()
	return math.Sqrt(float64(r-rr)*float64(r-rr) +
		float64(g-rg)*float64(g-rg) +
		float64(b-rb)*float64(b-rb))
}
