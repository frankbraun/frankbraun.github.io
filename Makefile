all: build-page
	./build-page

build-page: build-page.go
	go build -v $<

clean:
	rm -f build-page
