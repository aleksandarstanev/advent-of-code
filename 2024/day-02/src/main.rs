use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<Vec<i32>>> {
    let mut input = Vec::new();

    for line in reader.lines() {
        let line = line?;

        let iter = line.split_whitespace();

        let mut vec = Vec::new();
        for i in iter {
            vec.push(i.parse().unwrap());
        }

        input.push(vec);
    }

    return Ok(input);
}

fn check_increasing(vec: Vec<i32>, skip_idx: usize) -> bool {
    let mut prev = -1;
    for i in 0..vec.len() {
        if i == skip_idx {
            continue;
        }

        if prev != -1 && (vec[i] <= prev || vec[i] - prev > 3) {
            return false;
        }

        prev = vec[i];
    }

    return true;
}

fn check_decreasing(vec: Vec<i32>, skip_idx: usize) -> bool {
    let mut prev = -1;
    for i in 0..vec.len() {
        if i == skip_idx {
            continue;
        }

        if prev != -1 && (vec[i] >= prev || prev - vec[i] > 3) {
            return false;
        }

        prev = vec[i];
    }

    return true;
}

fn solve(reports: Vec<Vec<i32>>, can_remove: bool) -> i32 {
    let mut safe_levels = 0;

    for i in 0..reports.len() {
        let mut is_safe = false;
        if check_increasing(reports[i].clone(), usize::MAX)
            || check_decreasing(reports[i].clone(), usize::MAX)
        {
            is_safe = true;
        }

        if !can_remove {
            if is_safe {
                safe_levels += 1;
            }

            continue;
        }

        for skip in 0..reports[i].len() {
            if check_increasing(reports[i].clone(), skip)
                || check_decreasing(reports[i].clone(), skip)
            {
                is_safe = true;
            }
        }

        if is_safe {
            safe_levels += 1;
        }
    }

    return safe_levels;
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    let safe_reports = solve(input.clone(), true);

    println!("{}", safe_reports);

    Ok(())
}
