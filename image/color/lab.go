package color

import (
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"math"
)

// An implementation of the go color.Model interface
type LabColorModel struct {
}

func (cm *LabColorModel) Convert(in color.Color) color.Color {
	if c, ok := in.(*LabColor); ok {
		return c
	}

	r, g, b, _ := in.RGBA()
	rgba := colorful.Color{float64(r) / 65535.0, float64(g) / 65535.0, float64(b) / 65535.0}
	ll, aa, bb := rgba.Lab()
	return &LabColor{
		L: ll,
		A: aa,
		B: bb,
	}
}

// An implementation of DistanceMetric for the LabColor model.
func (m *LabColorModel) Distance(from, to color.Color) float64 {
	fromLab := m.Convert(from).(*LabColor)
	toLab := m.Convert(to).(*LabColor)
	return math.Sqrt(sq(fromLab.L-toLab.L) + sq(fromLab.A-toLab.A) + sq(fromLab.B-toLab.B))
}

// Helper function to square two float64
func sq(n float64) float64 {
	return n * n
}

// A point in the Lab color space. Implements the go image.Color interface
type LabColor struct {
	L float64
	A float64
	B float64
}

// Convert to RGBA colorspace
func (c *LabColor) RGBA() (uint32, uint32, uint32, uint32) {
	return colorful.Lab(c.L, c.A, c.B).RGBA()
}
