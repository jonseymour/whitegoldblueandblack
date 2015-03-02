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
	output/random-16.png \
	output/random-32.png \
	output/distance-rgb.png \
	output/distance-lab.png \
	output/snippet-distance.png \
	output/gold-snippet.png \
	output/white-snippet.png \
	output/blacken.png \
	output/snippet-random.png \
	output/snippet-randomize-blocks.png \
	output/mix-16x16.png \
	output/mix-3x495.png \
	output/snippet-mix.png

output/original.png: input/original.jpg
	$(WGBB) --jpeg  < $<  > $@

output/random.png: output/original.png
	$(WGBB) --randomize-blocks  < $<  > $@

output/random-16.png: output/original.png
	$(WGBB) --randomize-blocks --width=16 --height=16  < $<  > $@

output/random-32.png: output/original.png
	$(WGBB) --randomize-blocks --width=32 --height=32  < $<  > $@

output/mix-16x16.png: output/original.png
	$(WGBB) --mix-blocks --width=16 --height=16  < $<  > $@

output/mix-3x495.png: output/original.png
	$(WGBB) --mix-blocks --height=3 --width=495  < $<  > $@

output/distance-rgb.png: output/original.png
	$(WGBB) --sort-by-distance  < $<  > $@

output/distance-lab.png: output/original.png
	$(WGBB) --lab --sort-by-distance  < $<  > $@

output/snippet-distance.png: input/snippet.jpg
	$(WGBB) --jpeg --sort-by-distance < $<  > $@

output/snippet-random.png: input/snippet.jpg
	$(WGBB) --jpeg --randomize --stride=$(STRIDE) < $<  > $@

output/snippet-randomize-blocks.png: input/snippet.jpg
	$(WGBB) --jpeg --randomize-blocks  < $<  > $@

output/snippet-mix.png: input/snippet.jpg
	$(WGBB) --jpeg --mix-blocks --height 3 --width 210  < $<  > $@

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