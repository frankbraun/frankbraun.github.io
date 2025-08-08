use comrak::{Options, markdown_to_html};
use std::fmt::Write;
use std::fs::{self, File};
use std::io::{BufRead, BufReader, Error, ErrorKind, Result};

const DOCTYPE: &str = r#"<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
"#;

const TWITTER_CARD: &str = r#"<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@thefrankbraun">
<meta name="twitter:creator" content="@thefrankbraun">
"#;

const CLOSE_HEADER: &str = r#"<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="/css/water.css">
<link rel="alternate" type="application/atom+xml" title="Frank Braun" href="/atom.xml">
<link rel="icon" href="/favicon.ico" sizes="any">
<link rel="icon" type="image/png" sizes="16x16"  href="/favicon-16x16.png">
<link rel="icon" type="image/png" sizes="32x32"  href="/favicon-32x32.png">
<link rel="apple-touch-icon" href="/apple-touch-icon.png">
</head>
<body>
<main>
"#;

const DONATION: &str = r#"<p>(<em>If you like my work, please consider <a href="/donate">making a donation</a>.</em>)</p>
"#;

const FOOTER: &str = r#"</main>
</body>
</html>
"#;

/*
const ATOM_HEADER: &str = r#"<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="en">
<title>Frank Braun</title>
<id>https://frankbraun.org/atom.xml</id>
<author>
  <name>Frank Braun</name>
</author>
<icon>https://frankbraun.org/favicon-32x32.png</icon>
<link rel="self" type="application/atom+xml" href="https://frankbraun.org/atom.xml"/>
<link rel="alternate" type="text/html" href="/"/>
"#;

const ATOM_FOOTER: &str = r#"</feed>
"#;
*/

fn read_first_line(filename: &str) -> Result<String> {
    let fp = File::open(filename)?;
    let mut lines = BufReader::new(fp).lines();
    let line = match lines.next() {
        Some(res) => res?,
        None => return Err(Error::new(ErrorKind::UnexpectedEof, "file is empty")),
    };
    Ok(line.trim_start_matches("# ").to_string())
}

fn get_card(_filename: &str, _title: &str) -> Result<String> {
    let s = String::from(TWITTER_CARD);

/*
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
    */

    Ok(s)
}

fn build_page(filename: &str) -> Result<()> {
    let input = fs::read_to_string(filename)?;
    let mut options = Options::default();
    options.extension.header_ids = Some(String::new());
    let out = markdown_to_html(&input, &options);
    let output = filename.replace(".md", ".html");
    let title = read_first_line(filename)?;
    let mut s = String::from(DOCTYPE);
    write!(s, "<title> {title}")?;
    if filename != "index.md" {
        s.push_str(" | Frank Braun");

    }
    let card = get_card(filename, &title)?;
    s.push_str(&card);
    s.push_str("</title>\n");
    s.push_str(CLOSE_HEADER);
    s.push_str(&out);
    s.push_str(DONATION);
    s.push_str(FOOTER);
    fs::write(output, s)?;
    Ok(())
}

fn build_pages() -> Result<()> {
    build_page("index.md")
}

fn main() {
    if let Err(e) = build_pages() {
        eprintln!("error: {e:?}");
    }
}
