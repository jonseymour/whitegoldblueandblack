package main

import (
	distance "github.com/jonseymour/whitegoldblueandblack/image/color"
	"image"
	"image/color"
	"sort"
)

type distanceSort struct {
	img         image.Image
	ref         color.Color
	permutation []image.Point
	metric      distance.DistanceMetric
}

func newDistanceSort(img image.Image, ref color.Color, metric distance.DistanceMetric) *distanceSort {
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
	w := &distanceSort{
		img:         img,
		ref:         ref,
		permutation: permutation,
		metric:      metric,
	}
	sort.Sort(w)
	return w
}

func (w *distanceSort) Len() int {
	bounds := w.img.Bounds()
	return (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
}

func (w *distanceSort) Less(i, j int) bool {
	iImg := w.permutation[i]
	jImg := w.permutation[j]
	iColor := w.img.At(iImg.X, iImg.Y)
	jColor := w.img.At(jImg.X, jImg.Y)
	return w.metric(iColor, w.ref) < w.metric(jColor, w.ref)
}

func (w *distanceSort) Swap(i, j int) {
	tmp := w.permutation[i]
	w.permutation[i] = w.permutation[j]
	w.permutation[j] = tmp
}

// sorts all the pixels of an image by their RGB distance, then use
// a zig-zag sort to permute the pixels so that the darkest pixels
// are at the top-left and the distance pixels at the bottom right.
func sortByDistance(img image.Image, ref color.Color, metric distance.DistanceMetric) [][]image.Point {
	bounds := img.Bounds()
	lenX := (bounds.Max.X - bounds.Min.X)
	lenY := (bounds.Max.Y - bounds.Min.Y)
	w := newDistanceSort(img, ref, metric)
	z := newZigZagSort(lenX, lenY)

	permutation := make([][]image.Point, lenX)
	for i, _ := range permutation {
		permutation[i] = make([]image.Point, lenY)
	}
	for i, p := range w.permutation {
		permutation[z.permutation[i].X][z.permutation[i].Y] = p
	}
	return permutation
}
