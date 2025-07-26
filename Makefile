ASCIIART_SRC:=$(wildcard asciiart/*.txt)
ASCIIART_DST:=$(ASCIIART_SRC:%.txt=%.svg)

all: $(ASCIIART_DST) build-page
	./build-page

build-page: build-page.go
	go build -v $<

asciiart/%.svg: asciiart/%.txt
	svgbob -o $@ $<

clean:
	rm -f build-page

.PHONY: clean
