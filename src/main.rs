use clap::Parser;
use comrak::{Options, markdown_to_html};
use std::fmt::Write;
use std::fs::{self, File};
use std::io::{BufRead, BufReader, Error, ErrorKind, Result};
use walkdir::WalkDir;

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

const DONATION_OR_SUBSCRIBE: &str = r#"<p>(<em>If you like my work, please consider <a href="/donate">making a donation</a> or <a href="/#subscribe">subscribing</a>.</em>)</p>
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

fn truncate_string(s: &str, max_length: usize) -> String {
    if s.chars().count() <= max_length - 3 {
        return s.to_string();
    }
    let mut t: String = s.chars().take(max_length - 3).collect();
    t.push_str("...");
    t
}

fn get_description(filename: &str) -> Result<String> {
    let fp = File::open(filename)?;
    let mut lines = BufReader::new(fp).lines();
    loop {
        let line = match lines.next() {
            Some(res) => res?,
            None => return Err(Error::new(ErrorKind::UnexpectedEof, "file is empty")),
        };
        if line.starts_with("# ")
            || line.is_empty()
            || line.starts_with("*by")
            || line.starts_with("**")
            || line.starts_with("## ")
            || line.starts_with("(")
        {
            continue;
        }
        return Ok(truncate_string(&line, 150));
    }
}

fn get_card(filename: &str, title: &str) -> Result<String> {
    let mut s = String::from(TWITTER_CARD);
    let desc = get_description(filename)?;
    writeln!(s, "<meta property=\"og:title\" content=\"{title}\">").unwrap();
    writeln!(s, "<meta property=\"og:description\" content=\"{desc}\">").unwrap();
    writeln!(
        s,
        "<meta property=\"og:image\" content=\"https://frankbraun.org/img/frank-braun.png\">"
    )
    .unwrap();
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
    write!(s, "<title>{title}").unwrap();
    if filename != "index.md" {
        s.push_str(" | Frank Braun");
    }
    s.push_str("</title>\n");
    let card = get_card(filename, &title)?;
    s.push_str(&card);
    s.push_str(CLOSE_HEADER);
    s.push_str(&out);
    if filename != "donate/index.md" {
        if filename == "index.md" {
            s.push_str(DONATION);
        } else {
            s.push_str(DONATION_OR_SUBSCRIBE);
        }
    }
    s.push_str(FOOTER);
    fs::write(output, s)?;
    Ok(())
}

fn build_pages() -> Result<()> {
    for entry in WalkDir::new(".").into_iter().filter_map(|e| e.ok()) {
        if entry.file_type().is_file() {
            let path = entry.path().to_str().unwrap();
            if path != "./README.md" && path != "./CLAUDE.md" && path.ends_with(".md") {
                build_page(path)?;
            }
        }
    }
    Ok(())
}

fn write_feed() -> Result<()> {
    Ok(())
}

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[arg(long)]
    atom: bool,
}

fn main() {
    let args = Args::parse();
    if let Err(e) = build_pages() {
        return eprintln!("error: {e:?}");
    }
    if args.atom {
        if let Err(e) = write_feed() {
            eprintln!("error: {e:?}");
        }
    }
}
