package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

const doctype = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
`

const header = `<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="https://frankbraun.org/css/water.css">
</head>
<body>
<main>
`

const donation = `<p>(<em>If you like my work, please consider <a href="https://frankbraun.org/donate">making a donation</a>.</em>)</p>`

const footer = `</main>
</body>
</html>
`

func readFirstLine(filename string) (string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	if scanner.Scan() {
		l := scanner.Text()
		return strings.TrimPrefix(l, "# "), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("file '%s' is empty\n")
}

func buildPage() error {
	md := goldmark.New(goldmark.WithParserOptions(parser.WithAutoHeadingID()))
	// find all .md files in tree
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			// only consider .md files that are not README.md
			if path != "README.md" && strings.HasSuffix(path, ".md") {
				// read .md file in
				in, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				// convert it to HTML
				var out bytes.Buffer
				if err := md.Convert(in, &out); err != nil {
					return err
				}

				// create .html file
				fp, err := os.Create(strings.TrimSuffix(path, ".md") + ".html")
				if err != nil {
					return err
				}
				defer fp.Close()

				title, err := readFirstLine(path)
				if err != nil {
					return err
				}

				// write HTML to .html file
				if _, err := fp.WriteString(doctype); err != nil {
					return err
				}
				s := "<title>" + title
				if path != "index.md" {
					s += " | Frank Braun"
				}
				s += "</title>\n"
				if _, err := fp.WriteString(s); err != nil {
					return err
				}
				if _, err := fp.WriteString(header); err != nil {
					return err
				}
				if _, err := fp.Write(out.Bytes()); err != nil {
					return err
				}
				if path != "donate/index.md" {
					if _, err := fp.WriteString(donation); err != nil {
						return err
					}
				}
				if _, err := fp.WriteString(footer); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
	os.Exit(1)
}

func main() {
	if err := buildPage(); err != nil {
		fatal(err)
	}
}
