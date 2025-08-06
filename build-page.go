package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

const doctype = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
`
const twitterCard = `<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@thefrankbraun">
<meta name="twitter:creator" content="@thefrankbraun">
`

const closeHeader = `<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="/css/water.css">
<link rel="alternate" type="application/atom+xml" title="Frank Braun" href="/atom.xml">
<link rel="icon" href="/favicon.ico" sizes="any">
<link rel="icon" type="image/png" sizes="16x16"  href="/favicon-16x16.png">
<link rel="icon" type="image/png" sizes="32x32"  href="/favicon-32x32.png">
<link rel="apple-touch-icon" href="/apple-touch-icon.png">
</head>
<body>
<main>
`

const donation = `<p>(<em>If you like my work, please consider <a href="/donate">making a donation</a>.</em>)</p>
`
const donateOrSubscribe = `<p>(<em>If you like my work, please consider <a href="/donate">making a donation</a> or <a href="/#subscribe">subscribing</a>.</em>)</p>
`

const footer = `</main>
</body>
</html>
`

const atomHeader = `<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="en">
<title>Frank Braun</title>
<id>https://frankbraun.org/atom.xml</id>
<author>
  <name>Frank Braun</name>
</author>
<icon>https://frankbraun.org/favicon-32x32.png</icon>
<link rel="self" type="application/atom+xml" href="https://frankbraun.org/atom.xml"/>
<link rel="alternate" type="text/html" href="/"/>
`

const atomFooter = `</feed>
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
	return "", fmt.Errorf("file is empty: %s", filename)
}

func truncateString(s string, maxLength int) string {
	if len(s) <= maxLength-3 {
		return s
	}
	return s[:maxLength-3] + "..."
}

func getDescription(filename string) (string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "# ") ||
			l == "" ||
			strings.HasPrefix(l, "*by") ||
			strings.HasPrefix(l, "**") ||
			strings.HasPrefix(l, "## ") ||
			strings.HasPrefix(l, "(") {
			continue
		}
		return truncateString(l, 150), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no description found in: %s", filename)
}

func writeCard(fp *os.File, title, path string) error {
	desc, err := getDescription(path)
	if err != nil {
		return err
	}
	if _, err := fp.WriteString(twitterCard); err != nil {
		return err
	}
	fmt.Fprintf(fp, "<meta property=\"og:title\" content=\"%s\">\n", title)
	fmt.Fprintf(fp, "<meta property=\"og:description\" content=\"%s\">\n", desc)
	fmt.Fprintf(fp, "<meta property=\"og:image\" content=\"https://frankbraun.org/img/frank-braun.png\">\n")
	return nil
}

func buildPage() error {
	md := goldmark.New(goldmark.WithParserOptions(parser.WithAutoHeadingID()))
	// find all .md files in tree
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			// only consider .md files that are not README.md or CLAUDE.md
			if path != "README.md" && path != "CLAUDE.md" && strings.HasSuffix(path, ".md") {
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
				if err := writeCard(fp, title, path); err != nil {
					return err
				}
				if _, err := fp.WriteString(closeHeader); err != nil {
					return err
				}
				if _, err := fp.Write(out.Bytes()); err != nil {
					return err
				}
				if path != "donate/index.md" {
					if path == "index.md" {
						if _, err := fp.WriteString(donation); err != nil {
							return err
						}
					} else {
						if _, err := fp.WriteString(donateOrSubscribe); err != nil {
							return err
						}
					}
				}
				if _, err := fp.WriteString(footer); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func fileTime(filename string) (time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return time.Now(), err
	}
	return fi.ModTime(), nil
}

func isNewer(fileA, fileB string) (bool, error) {
	tA, err := fileTime(fileA)
	if err != nil {
		return false, err
	}
	tB, err := fileTime(fileB)
	if err != nil {
		return false, err
	}
	return tA.After(tB), nil
}

func writeFeedEntry(fp *os.File, desc string) error {
	re := regexp.MustCompile(`\[(.*?)\]\((.*?)\) \((.*?)\)`)
	matches := re.FindStringSubmatch(desc)
	if len(matches) != 4 {
		return fmt.Errorf("cannot match post line: %s", desc)
	}
	title := matches[1]
	link := matches[2]
	date := matches[3]
	desc, err := getDescription(strings.TrimPrefix(link, "/") + "/index.md")
	if err != nil {
		return err
	}
	s := "<entry>\n"
	s += "  <title>" + title + "</title>\n"
	s += "  <id>https://frankbraun.org" + link + "</id>\n"
	s += "  <link rel=\"alternate\" type=\"text/html\" href=\"https://frankbraun.org" + link + "\"/>\n"
	s += "  <updated>" + date + "T00:00:00Z" + "</updated>\n"
	s += "  <summary>" + desc + "</summary>\n"
	s += "</entry>\n"
	if _, err := fp.WriteString(s); err != nil {
		return err
	}
	return nil
}

func writeFeedEntries(fp *os.File) error {
	ip, err := os.Open("index.md")
	if err != nil {
		return err
	}
	defer ip.Close()
	scanner := bufio.NewScanner(ip)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "## Posts" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "#") {
			break
		}
		if strings.HasPrefix(l, "- ") {
			err := writeFeedEntry(fp, strings.TrimPrefix(l, "- "))
			if err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

func writeFeed() error {
	ok, err := isNewer("index.md", "atom.xml")
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	fp, err := os.Create("atom.xml")
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := fp.WriteString(atomHeader); err != nil {
		return err
	}
	s := "<updated>"
	s += time.Now().UTC().Format(time.RFC3339)
	s += "</updated>\n"
	if _, err := fp.WriteString(s); err != nil {
		return err
	}
	if err := writeFeedEntries(fp); err != nil {
		return err
	}
	if _, err := fp.WriteString(atomFooter); err != nil {
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
	if err := writeFeed(); err != nil {
		fatal(err)
	}
}
