# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a custom static site generator for frankbraun.org, a personal website focused on cryptocurrency, privacy technology, and programming essays. The generator is written in Go and converts Markdown files to HTML pages.

## Build Commands

```bash
# Build the entire site (compiles generator and runs it)
make

# Clean build artifacts
make clean

# Build just the Go generator without running it
go build -v build-page.go

# Run the generator after building
./build-page
```

## Architecture

### Core Generator (`build-page.go`)
The site generator performs these key functions:
1. Walks the directory tree to find `index.md` files
2. Converts Markdown to HTML using goldmark parser
3. Applies consistent HTML templating with:
   - Meta tags for SEO
   - Twitter Card and OpenGraph metadata
   - Proper HTML5 structure with navigation
4. Generates `atom.xml` feed from posts listed in the root `index.md`
5. Handles special case for ASCII art SVG generation via svgbob

### Content Structure
- Each topic/essay lives in its own directory with an `index.md` file
- The generator produces `index.html` in the same directory
- ASCII art files in `/asciiart/*.txt` are converted to SVG files
- Static assets (CSS, images) are served as-is

### Key Dependencies
- `github.com/yuin/goldmark` - Markdown parser
- `svgbob` - ASCII to SVG converter (must be installed separately)

## Development Notes

When modifying the site generator:
- The main logic is in `build-page.go`
- Template changes affect all generated pages
- New content directories should follow the pattern of existing ones (contain `index.md`)
- ASCII art files should be placed in `/asciiart/` with `.txt` extension

When adding new content:
- Create a new directory for major topics
- Add an `index.md` file with the content
- Run `make` to generate the HTML
- For blog posts, add an entry to the root `index.md` to include it in the Atom feed