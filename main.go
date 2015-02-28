// whitegoldblueandblack - (c) Jon Seymour 2015
//
// Some tools to help explore the intriguing case of the white and gold (or blue and black) dress.
//
// See the README for more details.
//

package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {

	stride := 1
	randomize := false
	brightness := false
	readJpeg := false
	permute := false

	flag.BoolVar(&brightness, "sort-by-brightness", false, "sort the image by brighness using blocks of width stride")
	flag.IntVar(&stride, "stride", 1, "Number of pixels to shift.")
	flag.BoolVar(&randomize, "randomize", false, "randomly sort the rows and colums of the image.")
	flag.BoolVar(&readJpeg, "jpeg", false, "The input is a jpeg rather than png.")
	flag.Parse()

	var img image.Image
	var err error

	processImage := readJpeg || randomize || brightness

	if processImage {
		if readJpeg {
			img, err = jpeg.Decode(os.Stdin)
		} else {
			img, err = png.Decode(os.Stdin)
		}
	}

	var permutation [][]image.Point

	if randomize {
		permutation = randomizeRowsAndColumns(img, stride)
		permute = true
	} else if brightness {
		permutation = sortByBrightness(img)
		permute = true
	} else if readJpeg {
		// just fallthrough
	} else {
		fmt.Fprintf(os.Stderr, "whitegoldblankandblue - (c) Jon Seymour 2015\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "usage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	bounds := img.Bounds()
	outbounds := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: bounds.Max.X - bounds.Min.X,
			Y: bounds.Max.Y - bounds.Min.Y,
		},
	}
	out := image.NewRGBA(outbounds)

	for x := outbounds.Min.X; x < outbounds.Max.X; x++ {
		for y := outbounds.Min.Y; y < outbounds.Max.Y; y++ {
			if permute {
				out.Set(x, y, img.At(permutation[x][y].X, permutation[x][y].Y))
			} else {
				out.Set(x, y, img.At(x, y))
			}
		}
	}

	err = png.Encode(os.Stdout, out)
	_ = err
}
