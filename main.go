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
	"image/color"
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
	distance := false
	readJpeg := false
	permute := false
	runColorize := false
	minPercentile := 0
	maxPercentile := 50
	colorizeProbability := 1.0
	colorHex := "#000000"
	refColorHex := "#000000"

	flag.BoolVar(&randomize, "randomize", false, "randomly sort the rows and colums of the image using blocks of stride pixels dimension.")
	flag.IntVar(&stride, "stride", 1, "Size of the block used for randomizing.")
	flag.BoolVar(&distance, "sort-by-distance", false, "Sort the image by color space distance.")
	flag.BoolVar(&runColorize, "colorize", false, "colorize pixels with a distance between --min-percentile and --max-percentile of --ref-color.")
	flag.BoolVar(&readJpeg, "jpeg", false, "The input is a jpeg rather than png.")
	flag.IntVar(&minPercentile, "min-percentile", 0, "The minimum distance percentile for colorizing.")
	flag.IntVar(&maxPercentile, "max-percentile", 50, "The maximum distance percentile for colorizing.")
	flag.Float64Var(&colorizeProbability, "colorize-prob", 1.0, "The probability of colorizing.")
	flag.StringVar(&colorHex, "color", "#000000", "The color to use for colorizing.")
	flag.StringVar(&refColorHex, "ref-color", "#000000", "The reference color to use for distance sorting.")
	flag.Parse()

	var img image.Image
	var aColor, aRefColor color.Color
	var err error

	processImage := readJpeg || randomize || distance || runColorize

	if processImage {
		if readJpeg {
			img, err = jpeg.Decode(os.Stdin)
		} else {
			img, err = png.Decode(os.Stdin)
		}
	}

	aColor, err = colorful.Hex(colorHex)
	if err != nil {
		die(err.Error())
	}

	aRefColor, err = colorful.Hex(refColorHex)
	if err != nil {
		die(err.Error())
	}

	if runColorize {
		transform := &colorize{
			minPercentile:       minPercentile,
			maxPercentile:       maxPercentile,
			colorizeProbability: colorizeProbability,
			color:               aColor,
			refColor:            aRefColor,
		}
		img = transform.transform(img)
	}

	var permutation [][]image.Point

	if randomize {
		permutation = randomizeRowsAndColumns(img, stride)
		permute = true
	} else if distance {
		permutation = sortByDistance(img, aRefColor)
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
