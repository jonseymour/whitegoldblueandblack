#NAME
whitegoldblueandblack - tools to explore the great #whiteandgold #blueandblack mystery that erupted on the Internet in late February 2015.

#SYNOPSIS
	wgbb {option...} < image > image.png

	where option is one or more of:
		[--jpeg]
		[--color {hex-color}]
		[--ref-color {hex-color}]
		[--sort-by-distance]
		[--randomize [--stride=n]]
		[--colorize [--min-percentile=n] [--max-percentile=n] [--colorize-prob=p]]
		[--rgb]
		[--lab]

	stdin is a PNG or (--jpeg: JPEG) image. stdout is a PNG image.

#OPTIONS

##--jpeg

The input is a JPEG image, otherwise the input is a PNG image.

##--color

The color to be used for colorizing. Defaults to #000000

##--ref-color

The color to be used as the reference color for distance calculations during (--sort-by-distance) or colorization (--colorize). Defaults to #000000.

##--sort-by-distance

The pixels of the image are sorted in color space distance order - pixels closest to the reference color are sorted top-left, pixels furthest from the reference color are sorted bottom-right.

##--randomize

Columns and rows of {stride} pixel width of the input image are randomly permuted.

##--colorize

Pixels within the specified percentiles (--min-percentile, --max-percentile) of distance are replaced with the specified color (--color) with a specified probability (--colorize-prob).

##--rgb

Perform distance calculations in the RGB color space. This is the default option.

##--lab

Perform distance calculations in the [CIE-Lab](http://en.wikipedia.org/wiki/Lab_color_space) color space.

#BUILDING

Install go,
	make

#EXAMPLES

Here is the original image:

![original](input/original.jpg)

Here is the same image randomized with:

	wgbb --jpeg --randomize --stride=1 < input/original.jpeg > output/random.png

![randomized](doc/random.png)

Here is the same image sorted by distance:

	wgbb --jpeg --sort-by-distance < input/original.jpeg > output/distance.png

![sorted-by-distance](doc/distance.png)

Here is a sort by distance analysis and randomization done to a small snippet of the dress:

	wgbb --jpeg --sort-by-distance < input/snippet.jpeg > output/snippet-distance.png
	wgbb --jpeg --randomize< input/snippet.jpeg > output/snippet-random.png

![snippet](input/snippet.jpg) ![snippet-sorted-by-distance](doc/snippet-distance.png) ![snippet-random](doc/snippet-random.png)

Here is the same sort by distance analysis done to a snippet of the gold (or black) region near the neckline:

	wgbb --jpeg --sort-by-distance < input/gold-snippet.jpeg > output/gold-snippet.png

![gold-snippet-location](doc/gold-snippet-location.png)
<br/>
![gold-snippet](input/gold-snippet.jpg) ![gold-snippet-sorted-by-distance](doc/gold-snippet.png)

Here is the same sort by distance analysis done to a small snippet of the 'white (or blue)' region on the subject's left shoulder:

	wgbb --jpeg --sort-by-distance < input/white-snippet.jpeg > output/white-snippet.png

![white-snippet-location](doc/white-snippet-location.png)
<br/>
![white-snippet](input/white-snippet.jpg) ![white-snippet-sorted-by-distance](doc/white-snippet.png)

Here is what happens when the darkest 40% of pixels are replaced with black.

	wgbb --jpeg --colorize --min-percentile 0 --max-percentile 40 --color "#000000" --colorize-prob 1.0 < input/original.jpg  > output/blacken.png

![blackened](doc/blacken.png)

#REVISIONS

* 0.1.0 - development release