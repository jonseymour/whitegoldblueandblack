GO_FILES=main.go randomize.go zigzag.go brightness.go
WGBB=bin/wgbb

all: $(WGBB) output/random.png output/brightness.png output/snippet-brightness.png output/gold-snippet.png output/white-snippet.png

output/original.png: input/original.jpg
	$(WGBB) --jpeg  < $<  > $@

output/random.png: output/original.png
	$(WGBB) --randomize  < $<  > $@

output/brightness.png: output/original.png
	$(WGBB) --sort-by-brightness  < $<  > $@

output/snippet-brightness.png: input/snippet.jpg
	$(WGBB) --jpeg --sort-by-brightness < $<  > $@

output/gold-snippet.png: input/gold-snippet.jpg
	$(WGBB) --jpeg --sort-by-brightness < $<  > $@

output/white-snippet.png: input/white-snippet.jpg
	$(WGBB) --jpeg --sort-by-brightness < $<  > $@

bin/wgbb: $(GO_FILES)
	go build -o bin/wgbb

clean:
	rm output/*