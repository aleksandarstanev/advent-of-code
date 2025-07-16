use std::collections::BTreeSet;
use std::fs::File;
use std::io::{self, BufRead};

// fn parse_input(mut reader: io::BufReader<File>) -> io::Result<String> {
// }

fn main() -> io::Result<()> {
    let file_path = "example.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    // let input = parse_input(reader)?;

    Ok(())
}
