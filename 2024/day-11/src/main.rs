use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<i64>> {
    let mut lines = reader.lines();
    let first_line = lines.next().unwrap()?;
    let numbers = first_line
        .split_whitespace()
        .map(|s| s.parse::<i64>().unwrap())
        .collect();

    Ok(numbers)
}

fn calculate_num_digits(num: i64) -> u32 {
    num.to_string().len() as u32
}

fn split_number(num: i64) -> (i64, i64) {
    let num_str = num.to_string();
    let first_half = num_str[..num_str.len() / 2].parse::<i64>().unwrap();
    let second_half = num_str[num_str.len() / 2..].parse::<i64>().unwrap();
    (first_half, second_half)
}

fn part_one(mut numbers: Vec<i64>) -> i32 {
    let iterations = 25;

    for _ in 0..iterations {
        let mut new_numbers: Vec<i64> = Vec::new();

        for number in numbers {
            if number == 0 {
                new_numbers.push(1);
            } else if calculate_num_digits(number) % 2 == 0 {
                let (first_half, second_half) = split_number(number);
                new_numbers.push(first_half);
                new_numbers.push(second_half);
            } else {
                new_numbers.push(number * 2024)
            }
        }

        numbers = new_numbers;
    }

    numbers.len() as i32
}

fn part_two(numbers: Vec<i64>) -> i64 {
    let iterations = 75;

    let mut num_count: HashMap<i64, i64> = HashMap::new();

    for number in numbers {
        *num_count.entry(number).or_insert(0) += 1;
    }

    for _ in 0..iterations {
        let mut new_count: HashMap<i64, i64> = HashMap::new();

        for (num, count) in num_count.iter() {
            if *num == 0 {
                *new_count.entry(1).or_insert(0) += *count;
            } else if calculate_num_digits(*num) % 2 == 0 {
                let (first_half, second_half) = split_number(*num);
                *new_count.entry(first_half).or_insert(0) += *count;
                *new_count.entry(second_half).or_insert(0) += *count;
            } else {
                *new_count.entry(num * 2024).or_insert(0) += *count;
            }
        }

        num_count = new_count;
    }

    let mut total_count = 0;
    for (_, count) in num_count.iter() {
        total_count += *count;
    }

    total_count
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    // let result = part_one(input.clone());

    // println!("Result: {}", result);

    let result = part_two(input.clone());

    println!("Result: {}", result);

    Ok(())
}

// 65601038650482 -> too low
