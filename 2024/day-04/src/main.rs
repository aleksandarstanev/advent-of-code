use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<String>> {
    let mut input = Vec::new();

    for line in reader.lines() {
        let line = line?;

        input.push(line);
    }

    return Ok(input);
}

fn part_one(input: Vec<String>) -> io::Result<i32> {
    let mut occurrences = 0;

    let rows = input.len();
    let cols = input[0].len();

    for i in 0..rows {
        let line = &input[i];

        for j in 0..cols - 3 {
            let c1 = line.chars().nth(j).unwrap();
            let c2 = line.chars().nth(j + 1).unwrap();
            let c3 = line.chars().nth(j + 2).unwrap();
            let c4 = line.chars().nth(j + 3).unwrap();

            if c1 == 'X' && c2 == 'M' && c3 == 'A' && c4 == 'S' {
                occurrences += 1;
            }

            if c1 == 'S' && c2 == 'A' && c3 == 'M' && c4 == 'X' {
                occurrences += 1;
            }
        }
    }

    for i in 0..cols {
        for j in 0..rows - 3 {
            let c1 = input[j].chars().nth(i).unwrap();
            let c2 = input[j + 1].chars().nth(i).unwrap();
            let c3 = input[j + 2].chars().nth(i).unwrap();
            let c4 = input[j + 3].chars().nth(i).unwrap();

            if c1 == 'X' && c2 == 'M' && c3 == 'A' && c4 == 'S' {
                occurrences += 1;
            }

            if c1 == 'S' && c2 == 'A' && c3 == 'M' && c4 == 'X' {
                occurrences += 1;
            }
        }
    }

    for i in 0..rows - 3 {
        for j in 0..cols - 3 {
            let c1 = input[i].chars().nth(j).unwrap();
            let c2 = input[i + 1].chars().nth(j + 1).unwrap();
            let c3 = input[i + 2].chars().nth(j + 2).unwrap();
            let c4 = input[i + 3].chars().nth(j + 3).unwrap();

            if c1 == 'X' && c2 == 'M' && c3 == 'A' && c4 == 'S' {
                occurrences += 1;
            }

            if c1 == 'S' && c2 == 'A' && c3 == 'M' && c4 == 'X' {
                occurrences += 1;
            }
        }
    }

    for i in 3..rows {
        for j in 0..cols - 3 {
            let c1 = input[i].chars().nth(j).unwrap();
            let c2 = input[i - 1].chars().nth(j + 1).unwrap();
            let c3 = input[i - 2].chars().nth(j + 2).unwrap();
            let c4 = input[i - 3].chars().nth(j + 3).unwrap();

            if c1 == 'X' && c2 == 'M' && c3 == 'A' && c4 == 'S' {
                occurrences += 1;
            }

            if c1 == 'S' && c2 == 'A' && c3 == 'M' && c4 == 'X' {
                occurrences += 1;
            }
        }
    }

    return Ok(occurrences);
}

fn part_two(input: Vec<String>) -> io::Result<i32> {
    let mut occurrences = 0;

    let rows = input.len();
    let cols = input[0].len();

    for i in 1..rows - 1 {
        for j in 1..cols - 1 {
            let c11 = input[i - 1].chars().nth(j - 1).unwrap();
            let c13 = input[i - 1].chars().nth(j + 1).unwrap();
            let c22 = input[i].chars().nth(j).unwrap();
            let c31 = input[i + 1].chars().nth(j - 1).unwrap();
            let c33 = input[i + 1].chars().nth(j + 1).unwrap();

            if c22 == 'A'
                && (c11 == 'M' && c33 == 'S' || c11 == 'S' && c33 == 'M')
                && (c13 == 'M' && c31 == 'S' || c13 == 'S' && c31 == 'M')
            {
                occurrences += 1;
            }
        }
    }

    return Ok(occurrences);
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    let occurrences = part_one(input.clone())?;

    println!("Part one: {}", occurrences);

    let occurrences = part_two(input.clone())?;

    println!("Part two: {}", occurrences);

    Ok(())
}
