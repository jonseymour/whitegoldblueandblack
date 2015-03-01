GO_FILES=\
	./colorful.go \
	./colorize.go \
	./distance.go \
	./main.go \
	./randomize.go \
	./util.go \
	./zigzag.go \
	./image/color/distance.go \
	./image/color/lab.go \
	./image/lab.go \
	./image/util/set.go

WGBB=bin/wgbb
MIN_PERCENTILE=0
MAX_PERCENTILE=40
COLORIZE_PROB=1.0
STRIDE=1


all: $(WGBB) \
	output/random.png \
	output/distance.png \
	output/snippet-distance.png \
	output/gold-snippet.png \
	output/white-snippet.png \
	output/blacken.png \
	output/snippet-random.png

output/original.png: input/original.jpg
	$(WGBB) --jpeg  < $<  > $@

output/random.png: output/original.png
	$(WGBB) --randomize  < $<  > $@

output/distance.png: output/original.png
	$(WGBB) --sort-by-distance  < $<  > $@

output/snippet-distance.png: input/snippet.jpg
	$(WGBB) --jpeg --sort-by-distance < $<  > $@

output/snippet-random.png: input/snippet.jpg
	$(WGBB) --jpeg --randomize --stride=$(STRIDE) < $<  > $@

output/gold-snippet.png: input/gold-snippet.jpg
	$(WGBB) --jpeg --sort-by-distance < $<  > $@

output/white-snippet.png: input/white-snippet.jpg
	$(WGBB) --jpeg --sort-by-distance < $<  > $@

output/blacken.png: input/original.jpg
	$(WGBB) --jpeg --colorize --min-percentile $(MIN_PERCENTILE) --max-percentile $(MAX_PERCENTILE) --colorize $(COLORIZE_PROB) < $<  > $@

bin/wgbb: $(GO_FILES)
	go get -d .
	go build -o bin/wgbb

clean:
	rm output/*