package main

import (
	"image"
	"math"
	"math/rand"
	"time"
)

// permute blocks of a slice stride elements long
func randomPermuteSlice(in []int, stride int) []int {
	// t, _ := time.Parse("2006-01-02 15:04:05", "2015-02-28 16:04:05")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]int, len(in))
	for i, e := range in {
		out[i] = e
	}

	strides := len(out)/stride - 1

	for i := 0; i < strides; i++ {
		j := int(r.Int31n(int32(strides-i))) + i
		for k := 0; k < stride; k++ {
			t := out[j*stride+k]
			out[j*stride+k] = out[i*stride+k]
			out[i*stride+k] = t
		}
	}

	return out
}

// Answer a permutation which will permute the rows and columns of an image
// randomly.
func randomizeRowsAndColumns(img image.Image, stride int) [][]image.Point {

	lenX := (img.Bounds().Max.X - img.Bounds().Min.X)
	lenY := (img.Bounds().Max.Y - img.Bounds().Min.Y)
	dim := int(math.Sqrt(float64(lenX * lenY)))

	permutation := make([][]image.Point, dim)
	for i, _ := range permutation {
		permutation[i] = make([]image.Point, dim)
	}

	permuteX := make([]int, lenX)
	permuteY := make([]int, lenY)

	for i, _ := range permuteX {
		permuteX[i] = i + img.Bounds().Min.X
	}

	for j, _ := range permuteY {
		permuteY[j] = j + img.Bounds().Min.Y
	}

	permuteX = randomPermuteSlice(permuteX, stride)
	permuteY = randomPermuteSlice(permuteY, stride)
	limit := dim * dim

	for i, _ := range permuteX {
		for j, _ := range permuteY {
			offset := i*lenY + j
			if offset == limit {
				break
			}
			ii := offset / dim
			jj := offset % dim
			permutation[ii][jj] = image.Point{X: permuteX[i], Y: permuteY[j]}
		}
	}
	return permutation
}

// randomize blocks - moves blocks of size width and height
func randomizeBlocks(img image.Image, width, height int) [][]image.Point {
	lenX := (img.Bounds().Max.X - img.Bounds().Min.X)
	lenY := (img.Bounds().Max.Y - img.Bounds().Min.Y)

	xBlocks := lenX / width
	yBlocks := lenY / height
	// normalize to whole blocks
	outLenX := xBlocks * width
	outLenY := yBlocks * height

	blocks := make([]int, xBlocks*yBlocks)
	for i, _ := range blocks {
		blocks[i] = i
	}
	blocks = randomPermuteSlice(blocks, 1)
	out := make([][]image.Point, outLenX)
	for i, _ := range out {
		out[i] = make([]image.Point, outLenY)
	}

	toSource := func(b int) (int, int) {
		bX := b % xBlocks
		bY := b / xBlocks

		bbX := bX * width
		bbY := bY * height

		return bbX, bbY
	}

	for i, b := range blocks {
		fx, fy := toSource(b)
		tx, ty := toSource(i)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				out[tx+x][ty+y] = image.Point{X: fx + x, Y: fy + y}
			}
		}
	}
	return out
}

func mixBlocks(img image.Image, width, height int) [][]image.Point {
	lenX := (img.Bounds().Max.X - img.Bounds().Min.X)
	lenY := (img.Bounds().Max.Y - img.Bounds().Min.Y)

	xBlocks := lenX / width
	yBlocks := lenY / height
	// normalize to whole blocks
	outLenX := xBlocks * width
	outLenY := yBlocks * height

	out := make([][]image.Point, outLenX)
	for i, _ := range out {
		out[i] = make([]image.Point, outLenY)
	}

	toSource := func(b int) (int, int) {
		bX := b % xBlocks
		bY := b / xBlocks

		bbX := bX * width
		bbY := bY * height

		return bbX, bbY
	}

	toXY := func(x int) (int, int) {
		return x % width, x / width
	}

	block := make([]int, width*height)
	for i, _ := range block {
		block[i] = i
	}

	for i := 0; i < xBlocks*yBlocks; i++ {
		block = randomPermuteSlice(block, 1)
		bX, bY := toSource(i)
		for i, b := range block {
			tx, ty := toXY(i)
			fx, fy := toXY(b)
			out[bX+tx][bY+ty] = image.Point{X: bX + fx, Y: bY + fy}
		}
	}
	return out
}
