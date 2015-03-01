package image

import (
	lab "github.com/jonseymour/whitegoldblueandblack/image/color"
	"image"
	"image/color"
)

// A LabImage stores its pixels as values in the Lab colorspace which is a colorspace
// which produces human-centric distance metrics
type LabImage struct {
	bounds image.Rectangle
	l      []float64
	a      []float64
	b      []float64
	minX   int
	minY   int
	dimX   int
	dimY   int
}

// Answers a new LabImage which is optimized for distance calculations.
func NewLabImage(bounds image.Rectangle) *LabImage {
	minX := bounds.Min.X
	minY := bounds.Min.Y
	dimX := bounds.Max.X - minX
	dimY := bounds.Max.Y - minY
	pixels := dimX * dimY
	return &LabImage{
		bounds: bounds,
		l:      make([]float64, pixels),
		a:      make([]float64, pixels),
		b:      make([]float64, pixels),
		minX:   minX,
		minY:   minY,
		dimX:   dimX,
		dimY:   dimY,
	}
}

// Answer the Lab color model.
func (i *LabImage) ColorModel() color.Model {
	return &lab.LabColorModel{}
}

// Answer the bounds of the image.
func (i *LabImage) Bounds() image.Rectangle {
	return i.bounds
}

// Answer the Lab color of the image at the specified offsets.
func (i *LabImage) At(x, y int) color.Color {
	offset := (x-i.minX)*i.dimY + (y - i.minY)
	return &lab.LabColor{
		L: i.l[offset],
		A: i.a[offset],
		B: i.b[offset],
	}
}

// Set the Lab color of the image at the specified offets.
func (i *LabImage) Set(x, y int, c color.Color) {
	lc := i.ColorModel().Convert(c).(*lab.LabColor)
	offset := (x-i.minX)*i.dimY + (y - i.minY)
	i.l[offset] = lc.L
	i.a[offset] = lc.A
	i.b[offset] = lc.B
}
