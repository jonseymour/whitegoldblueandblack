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
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func die(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}

func main() {

	stride := 1
	randomize := false
	brightness := false
	readJpeg := false
	permute := false
	runColorize := false
	minPercentile := 0
	maxPercentile := 50
	colorizeProbability := 1.0
	color := "#000000"

	flag.BoolVar(&brightness, "sort-by-brightness", false, "sort the image by brighness using blocks of width stride")
	flag.IntVar(&stride, "stride", 1, "Number of pixels to shift.")
	flag.BoolVar(&randomize, "randomize", false, "randomly sort the rows and colums of the image.")
	flag.BoolVar(&runColorize, "colorize", false, "randomly sort the rows and colums of the image.")
	flag.BoolVar(&readJpeg, "jpeg", false, "The input is a jpeg rather than png.")
	flag.IntVar(&minPercentile, "min-percentile", 0, "The minimum percentile for colorizing.")
	flag.IntVar(&maxPercentile, "max-percentile", 50, "The max percentile for colorizing.")
	flag.Float64Var(&colorizeProbability, "colorize-prob", 1.0, "The probability of colorizing.")
	flag.StringVar(&color, "color", "#000000", "The color to use for colorizing.")
	flag.Parse()

	var img image.Image
	var err error

	processImage := readJpeg || randomize || brightness || runColorize

	if processImage {
		if readJpeg {
			img, err = jpeg.Decode(os.Stdin)
		} else {
			img, err = png.Decode(os.Stdin)
		}
	}

	if runColorize {
		if theColor, err := colorful.Hex(color); err != nil {
			die(err.Error())
		} else {
			transform := &colorize{
				minPercentile:       minPercentile,
				maxPercentile:       maxPercentile,
				colorizeProbability: colorizeProbability,
				color:               theColor,
			}
			img = transform.transform(img)
		}
	}

	var permutation [][]image.Point

	if randomize {
		permutation = randomizeRowsAndColumns(img, stride)
		permute = true
	} else if brightness {
		permutation = sortByBrightness(img)
		permute = true
	} else if processImage {
		// just fallthrough
	} else {
		fmt.Fprintf(os.Stderr, "whitegoldblankandblue - (c) Jon Seymour 2015\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "usage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	out := sameSize(img)
	outbounds := out.Bounds()

	for x := 0; x < outbounds.Max.X; x++ {
		for y := 0; y < outbounds.Max.Y; y++ {
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
