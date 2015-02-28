GO_FILES=main.go randomize.go zigzag.go brightness.go util.go colorize.go
WGBB=bin/wgbb
MIN_PERCENTILE=0
MAX_PERCENTILE=40
BLACKEN_PROB=1.0
STRIDE=1


all: $(WGBB) output/random.png output/brightness.png output/snippet-brightness.png output/gold-snippet.png output/white-snippet.png output/blacken.png output/snippet-random.png

output/original.png: input/original.jpg
	$(WGBB) --jpeg  < $<  > $@

output/random.png: output/original.png
	$(WGBB) --randomize  < $<  > $@

output/brightness.png: output/original.png
	$(WGBB) --sort-by-brightness  < $<  > $@

output/snippet-brightness.png: input/snippet.jpg
	$(WGBB) --jpeg --sort-by-brightness < $<  > $@

output/snippet-random.png: input/snippet.jpg
	$(WGBB) --jpeg --randomize --stride=$(STRIDE) < $<  > $@

output/gold-snippet.png: input/gold-snippet.jpg
	$(WGBB) --jpeg --sort-by-brightness < $<  > $@

output/white-snippet.png: input/white-snippet.jpg
	$(WGBB) --jpeg --sort-by-brightness < $<  > $@

output/blacken.png: input/original.jpg
	$(WGBB) --jpeg --blacken --min-percentile $(MIN_PERCENTILE) --max-percentile $(MAX_PERCENTILE) --blacken-prob $(BLACKEN_PROB) < $<  > $@

bin/wgbb: $(GO_FILES)
	go build -o bin/wgbb

clean:
	rm output/*