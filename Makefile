GO_FILES=\
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
	output/random-2.png \
	output/random-8.png \
	output/random-16.png \
	output/distance-rgb.png \
	output/distance-lab.png \
	output/snippet-distance.png \
	output/gold-snippet.png \
	output/white-snippet.png \
	output/blacken.png \
	output/snippet-random.png

output/original.png: input/original.jpg
	$(WGBB) --jpeg  < $<  > $@

output/random.png: output/original.png
	$(WGBB) --randomize  < $<  > $@

output/random-2.png: output/original.png
	$(WGBB) --randomize --stride=2  < $<  > $@

output/random-8.png: output/original.png
	$(WGBB) --randomize --stride=8  < $<  > $@

output/random-16.png: output/original.png
	$(WGBB) --randomize --stride=16  < $<  > $@

output/distance-rgb.png: output/original.png
	$(WGBB) --sort-by-distance  < $<  > $@

output/distance-lab.png: output/original.png
	$(WGBB) --lab --sort-by-distance  < $<  > $@

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

update-doc:
	cd doc && \
	for f in *.png; do \
		! test -f ../output/$$f || cp ../output/$$f .; \
	done

clean:
	rm output/*