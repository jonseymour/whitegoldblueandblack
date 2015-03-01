package main

import (
	"image"
	"image/color"
	"math"
	"sort"
)

type brightnessSort struct {
	img         image.Image
	ref         color.Color
	length      int
	minX        int
	minY        int
	lenX        int
	lenY        int
	permutation []image.Point
}

func newBrightnessSort(img image.Image, ref color.Color) *brightnessSort {
	bounds := img.Bounds()
	minX := bounds.Min.X
	minY := bounds.Min.Y
	lenX := bounds.Max.X - bounds.Min.X
	lenY := bounds.Max.Y - bounds.Min.Y
	permutation := make([]image.Point, lenX*lenY)
	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			permutation[i*lenY+j] = image.Point{X: minX + i, Y: minY + j}
		}
	}
	w := &brightnessSort{
		img:         img,
		ref:         ref,
		length:      lenY * lenX,
		lenX:        lenX,
		lenY:        lenY,
		minX:        minX,
		minY:        minY,
		permutation: permutation,
	}
	sort.Sort(w)
	return w
}

func brightness(c color.Color, ref color.Color) float64 {
	r, g, b, _ := c.RGBA()
	rr, rg, rb, _ := ref.RGBA()
	return math.Sqrt(float64(r-rr)*float64(r-rr) +
		float64(g-rg)*float64(g-rg) +
		float64(b-rb)*float64(b-rb))
}

func (w *brightnessSort) Len() int {
	return w.length
}

func (w *brightnessSort) Less(i, j int) bool {
	iImg := w.permutation[i]
	jImg := w.permutation[j]
	iColor := w.img.At(iImg.X, iImg.Y)
	jColor := w.img.At(jImg.X, jImg.Y)
	return brightness(iColor, w.ref) < brightness(jColor, w.ref)
}

func (w *brightnessSort) Swap(i, j int) {
	tmp := w.permutation[i]
	w.permutation[i] = w.permutation[j]
	w.permutation[j] = tmp
}

// sorts all the pixels of an image by their RGB brightness, then use
// a zig-zag sort to permute the pixels so that the darkest pixels
// are at the top-left and the brightness pixels at the bottom right.
func sortByBrightness(img image.Image, ref color.Color) [][]image.Point {
	w := newBrightnessSort(img, ref)
	z := newZigZagSort(w.lenX, w.lenY)

	permutation := make([][]image.Point, w.lenX)
	for i, _ := range permutation {
		permutation[i] = make([]image.Point, w.lenY)
	}
	for i, p := range w.permutation {
		permutation[z.permutation[i].X][z.permutation[i].Y] = p
	}
	return permutation
}
