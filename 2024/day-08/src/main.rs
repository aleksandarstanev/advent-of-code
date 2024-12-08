use std::collections::{HashMap, HashSet};
use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<String>> {
    reader.lines().collect()
}

fn is_antenna(char: char) -> bool {
    char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9'
}

fn get_positions(input: Vec<String>) -> HashMap<char, Vec<(usize, usize)>> {
    let mut positions: HashMap<char, Vec<(usize, usize)>> = HashMap::new();

    for (row, line) in input.iter().enumerate() {
        for (col, char) in line.chars().enumerate() {
            if is_antenna(char) {
                positions.entry(char).or_insert(Vec::new()).push((row, col));
            }
        }
    }

    positions
}

fn get_antidotes(r1: i32, c1: i32, r2: i32, c2: i32, rows: i32, cols: i32) -> Vec<(i32, i32)> {
    let r_diff = r2 - r1;
    let c_diff = c2 - c1;

    let mut antidotes = Vec::new();

    if r1 + r_diff < rows && r1 + r_diff >= 0 && c1 + c_diff < cols && c1 + c_diff >= 0 {
        antidotes.push((r1 + r_diff, c1 + c_diff));
    }

    if r1 - r_diff < rows && r1 - r_diff >= 0 && c1 - c_diff < cols && c1 - c_diff >= 0 {
        antidotes.push((r1 - r_diff, c1 - c_diff));
    }

    if r2 + r_diff < rows && r2 + r_diff >= 0 && c2 + c_diff < cols && c2 + c_diff >= 0 {
        antidotes.push((r2 + r_diff, c2 + c_diff));
    }

    if r2 - r_diff < rows && r2 - r_diff >= 0 && c2 - c_diff < cols && c2 - c_diff >= 0 {
        antidotes.push((r2 - r_diff, c2 - c_diff));
    }

    antidotes
}

fn get_antidotes_part_two(
    r1: i32,
    c1: i32,
    r2: i32,
    c2: i32,
    rows: i32,
    cols: i32,
) -> Vec<(i32, i32)> {
    let r_diff = r2 - r1;
    let c_diff = c2 - c1;

    let mut antidotes = Vec::new();

    let mut rr = r1;
    let mut cc = c1;

    while rr >= 0 && rr < rows && cc >= 0 && cc < cols {
        antidotes.push((rr, cc));

        rr += r_diff;
        cc += c_diff;
    }

    rr = r1;
    cc = c1;

    while rr >= 0 && rr < rows && cc >= 0 && cc < cols {
        antidotes.push((rr, cc));

        rr -= r_diff;
        cc -= c_diff;
    }

    rr = r2;
    cc = c2;

    while rr >= 0 && rr < rows && cc >= 0 && cc < cols {
        antidotes.push((rr, cc));

        rr += r_diff;
        cc += c_diff;
    }

    rr = r2;
    cc = c2;

    while rr >= 0 && rr < rows && cc >= 0 && cc < cols {
        antidotes.push((rr, cc));

        rr -= r_diff;
        cc -= c_diff;
    }

    antidotes
}

fn part_one(input: Vec<String>) -> io::Result<i32> {
    let rows = input.len();
    let cols = input[0].len();

    let antenna_positions = get_positions(input);

    let mut antinodes = HashSet::new();

    for (_, positions) in antenna_positions {
        for i in 0..positions.len() {
            for j in i + 1..positions.len() {
                let (r1, c1) = positions[i];
                let (r2, c2) = positions[j];

                let candidates = get_antidotes(
                    r1 as i32,
                    c1 as i32,
                    r2 as i32,
                    c2 as i32,
                    rows as i32,
                    cols as i32,
                );

                for (r, c) in candidates {
                    if (r == r1 as i32 && c == c1 as i32) || (r == r2 as i32 && c == c2 as i32) {
                        continue;
                    }

                    antinodes.insert((r, c));
                }
            }
        }
    }

    Ok(antinodes.len() as i32)
}

fn part_two(input: Vec<String>) -> io::Result<i32> {
    let rows = input.len();
    let cols = input[0].len();

    let antenna_positions = get_positions(input);

    let mut antinodes = HashSet::new();

    for (_, positions) in antenna_positions {
        for i in 0..positions.len() {
            for j in i + 1..positions.len() {
                let (r1, c1) = positions[i];
                let (r2, c2) = positions[j];

                let candidates = get_antidotes_part_two(
                    r1 as i32,
                    c1 as i32,
                    r2 as i32,
                    c2 as i32,
                    rows as i32,
                    cols as i32,
                );

                for (r, c) in candidates {
                    antinodes.insert((r, c));
                }
            }
        }
    }

    Ok(antinodes.len() as i32)
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    // let result = part_one(input)?;

    // println!("Result: {}", result);

    let result = part_two(input)?;

    println!("Result: {}", result);

    Ok(())
}
