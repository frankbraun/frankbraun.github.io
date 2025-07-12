package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
)

const header = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Frank Braun</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="css/water.css">
</head>
<body>
<main>
`

const footer = `</main>
</body>
</html>
`

func buildPage() error {
	in, err := os.ReadFile("index.md")
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := goldmark.Convert(in, &out); err != nil {
		return err
	}

	fp, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer fp.Close()

	if _, err := fp.WriteString(header); err != nil {
		return err
	}
	if _, err := fp.Write(out.Bytes()); err != nil {
		return err
	}
	if _, err := fp.WriteString(footer); err != nil {
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
