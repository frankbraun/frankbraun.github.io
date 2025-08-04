use std::fs;

fn build_page() -> Result<(), std::io::Error> {
    let input = fs::read_to_string("index.md")?;
    let out = markdown::to_html(&input);
    print!("{out}");
    Ok(())
}

fn main() {
    if let Err(e) = build_page() {
        eprintln!("Error: {e}");
    }
}
