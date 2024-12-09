use std::collections::BTreeSet;
use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(mut reader: io::BufReader<File>) -> io::Result<String> {
    let mut line = String::new();
    reader.read_line(&mut line)?;
    Ok(line)
}

fn get_layout(input: String) -> (Vec<i32>, BTreeSet<i32>, Vec<(i32, i32)>) {
    let mut layout = Vec::<i32>::new();
    let mut empty_slots = BTreeSet::<i32>::new();
    let mut file_ranges = Vec::<(i32, i32)>::new();

    let mut file_id = 0;
    for (i, c) in input.chars().enumerate() {
        let digit = c.to_digit(10).unwrap();

        if i % 2 == 0 {
            file_ranges.push((layout.len() as i32, layout.len() as i32 + digit as i32 - 1));

            for _ in 0..digit {
                layout.push(file_id);
            }

            file_id += 1;
        } else {
            for _ in 0..digit {
                empty_slots.insert(layout.len() as i32);
                layout.push(-1);
            }
        }
    }

    (layout, empty_slots, file_ranges)
}

fn calculate_checksum(layout: Vec<i32>) -> i64 {
    let mut result: i64 = 0;
    for i in 0..layout.len() {
        if layout[i] != -1 {
            result += (layout[i] as i64) * (i as i64);
        }
    }

    result
}

fn part_one(input: String) -> io::Result<i64> {
    let (mut layout, mut empty_slots, _) = get_layout(input);

    let mut idx_to_swap = layout.len() - 1;
    for _ in 0..empty_slots.len() {
        while layout[idx_to_swap] == -1 {
            idx_to_swap -= 1;
        }

        let first_empty = *empty_slots.first().unwrap();
        if first_empty >= idx_to_swap as i32 {
            break;
        }

        layout[*empty_slots.first().unwrap() as usize] = layout[idx_to_swap];
        layout[idx_to_swap] = -1;

        empty_slots.remove(&first_empty);

        idx_to_swap -= 1;
    }

    let result = calculate_checksum(layout);

    Ok(result)
}

fn part_two(input: String) -> io::Result<i64> {
    let (mut layout, mut empty_slots, file_ranges) = get_layout(input);

    for i in (0..file_ranges.len()).rev() {
        let (start, end) = file_ranges[i];
        let length = end - start + 1;

        let mut last_slot = -1;
        let mut consecutive_length = 0;
        for j in empty_slots.iter() {
            if last_slot == -1 || *j != last_slot + 1 {
                consecutive_length = 1;
            } else if *j == last_slot + 1 {
                consecutive_length += 1;
            }

            last_slot = *j;

            if consecutive_length == length {
                break;
            }
        }

        if consecutive_length == length && last_slot < end {
            for j in 0..length {
                layout[last_slot as usize] = layout[start as usize + j as usize];
                layout[start as usize + j as usize] = -1;

                empty_slots.remove(&last_slot);
                last_slot -= 1;
            }
        }
    }

    let result = calculate_checksum(layout);

    Ok(result)
}

fn main() -> io::Result<()> {
    let file_path = "example.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    // let result = part_one(input)?;

    // println!("Result: {}", result);

    let result = part_two(input)?;

    println!("Result: {}", result);

    Ok(())
}
