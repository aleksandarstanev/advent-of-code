use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(reader: io::BufReader<File>) -> io::Result<(Vec<i32>, Vec<i32>)> {
    let mut first_vec: Vec<i32> = Vec::new();
    let mut second_vec: Vec<i32> = Vec::new();

    for line in reader.lines() {
        let line = line?;

        let mut iter = line.split_whitespace();

        let first: i32 = iter.next().unwrap().parse().unwrap();
        let second: i32 = iter.next().unwrap().parse().unwrap();

        first_vec.push(first);
        second_vec.push(second);
    }

    return Ok((first_vec, second_vec));
}

fn first_part((mut first_vec, mut second_vec): (Vec<i32>, Vec<i32>)) -> io::Result<i32> {
    first_vec.sort();
    second_vec.sort();

    let mut total_distance = 0;
    for i in 0..first_vec.len() {
        total_distance += (second_vec[i] - first_vec[i]).abs();
    }

    return Ok(total_distance);
}

fn second_part((first_vec, second_vec): (Vec<i32>, Vec<i32>)) -> io::Result<i32> {
    let mut count_occurences = HashMap::new();

    for i in 0..second_vec.len() {
        let count = count_occurences.entry(second_vec[i]).or_insert(0);
        *count += 1;
    }

    let mut similarity_score = 0;
    for i in 0..first_vec.len() {
        let count = count_occurences.get(&first_vec[i]);

        if count.is_some() {
            similarity_score += first_vec[i] * count.unwrap();
        }
    }

    return Ok(similarity_score);
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let (first_vec, second_vec) = parse_input(reader)?;

    let total_distance = first_part((first_vec.clone(), second_vec.clone()))?;

    println!("First part result: {}", total_distance);

    let similarity_score = second_part((first_vec.clone(), second_vec.clone()))?;

    println!("Second part result: {}", similarity_score);

    Ok(())
}
