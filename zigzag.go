package main

import (
	"image"
	"sort"
)

type zigZagSort struct {
	lenX        int
	lenY        int
	permutation []image.Point
}

// Produces a walk through matrix of length X and length Y such that
// the time t(x,y) that x,y is visited is such that t(x,y) > t(k,j)
// where x+y > k+j for all x,y,k and j.
//
// Note that this is weaker than a true zig-zag since the order of traversal
// along a diagonal isn't stable.
func newZigZagSort(lenX, lenY int) *zigZagSort {
	z := &zigZagSort{
		lenX:        lenX,
		lenY:        lenY,
		permutation: make([]image.Point, lenX*lenY),
	}
	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			z.permutation[i*lenY+j] = image.Point{X: i, Y: j}
		}
	}
	sort.Sort(z)
	return z
}

func (z *zigZagSort) Len() int {
	return z.lenX * z.lenY
}

func (z *zigZagSort) Less(i, j int) bool {
	iCoord := z.permutation[i]
	jCoord := z.permutation[j]
	return iCoord.X+iCoord.Y < jCoord.X+jCoord.Y
}

func (z *zigZagSort) Swap(i, j int) {
	t := z.permutation[i]
	z.permutation[i] = z.permutation[j]
	z.permutation[j] = t
}
