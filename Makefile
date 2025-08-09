ASCIIART_SRC:=$(wildcard asciiart/*.txt)
ASCIIART_DST:=$(ASCIIART_SRC:%.txt=%.svg)

all: $(ASCIIART_DST) ./target/debug/build-page
	./target/debug/build-page

atom: $(ASCIIART_DST) ./target/debug/build-page
	./target/debug/build-page --atom

./target/debug/build-page: src/main.rs
	cargo build

asciiart/%.svg: asciiart/%.txt
	svgbob -o $@ $<

.PHONY: clean cleanup check fmt
clean:
	rm -f ./target/debug/build-page

cleanup:
	rm -rf ./target

check:
	cargo clippy

fmt:
	cargo fmt
