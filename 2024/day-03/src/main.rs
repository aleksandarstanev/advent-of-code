use std::fs::read_to_string;
use std::io::{self};

use regex::Regex;

fn part_one(input: String) -> i64 {
    let re = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();

    let mut result: i64 = 0;
    for cap in re.captures_iter(&input) {
        let first_number = &cap[1];
        let second_number = &cap[2];

        let first_number: i64 = first_number.parse().unwrap();
        let second_number: i64 = second_number.parse().unwrap();

        result += first_number * second_number;
    }

    result
}

fn part_two(input: String) -> i64 {
    let re = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)").unwrap();

    let mut should_add = true;
    let mut result: i64 = 0;
    for cap in re.captures_iter(&input) {
        if let Some(first_number) = cap.get(1) {
            if !should_add {
                continue;
            }

            let first_number: i64 = first_number.as_str().parse().unwrap();
            let second_number: i64 = cap[2].parse().unwrap();
            result += first_number * second_number;
        } else {
            if &cap[0] == "do()" {
                should_add = true;
            } else {
                should_add = false;
            }
        }
    }

    result
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let input = read_to_string(file_path)?;

    println!("Part one: {}", part_one(input.clone()));

    println!("Part two: {}", part_two(input.clone()));

    Ok(())
}
