NAME
====
whitegoldblueandblack - tools to explore the great #whiteandgold #blueandblack mystery that erupted on the Internet in late February 2015.

SYNOPSIS
========
	wgbb [--jpeg] [--sort-by-brightness]|[--randomize [--stride=n]] < image > image

BUILDING
========
	make

EXAMPLES
========

Here is the original image:

![original](input/original.jpg)

Here is the same image randomized with:

	wgbb --jpeg --randomize --stride=1 < input/original.jpeg > output/random.png

![randomized](doc/random.png)

Here is the same image sorted by brightness:

	wgbb --jpeg --sort-by-brightness < input/original.jpeg > output/brightness.png

![sorted-by-brightness](doc/brightness.png)

Here is the same sort by brightness analysis done to a small snippet of the dress:

	wgbb --jpeg --sort-by-brightness < input/snippet.jpeg > output/snippet-brightness.png

![snippet](input/snippet.jpg) ![snippet-sorted-by-brightness](doc/snippet-brightness.png)

Here is the same sort by brightness analysis done to a snippet of the gold region near the neckline:

	wgbb --jpeg --sort-by-brightness < input/gold-snippet.jpeg > output/gold-snippet.png

![gold-snippet-location](doc/gold-snippet-location.png)
<br/>
![gold-snippet](input/gold-snippet.jpg) ![gold-snippet-sorted-by-brightness](doc/gold-snippet.png)

Here is the same sort by brightness analysis done to a small snippet of the white region on the subject's left shoulder:

	wgbb --jpeg --sort-by-brightness < input/white-snippet.jpeg > output/white-snippet.png

![white-snippet-location](doc/white-snippet-location.png)
<br/>
![white-snippet](input/white-snippet.jpg) ![white-snippet-sorted-by-brightness](doc/white-snippet.png)
