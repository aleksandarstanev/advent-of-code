use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<String>> {
    let mut lines = Vec::<String>::new();

    for line in reader.lines() {
        lines.push(line?);
    }

    Ok(lines)
}

fn traverse(grid: &Vec<String>, row: usize, col: usize, visited: &mut Vec<Vec<bool>>) -> i32 {
    let rows = grid.len();
    let cols = grid[0].len();

    let current = grid[row].chars().nth(col).unwrap();

    visited[row][col] = true;

    if current == '9' {
        return 1;
    }

    let mut count = 0;

    if row > 0
        && !visited[row - 1][col]
        && grid[row - 1].chars().nth(col).unwrap() == (current as u8 + 1) as char
    {
        count += traverse(grid, row - 1, col, visited);
    }

    if row < rows - 1
        && !visited[row + 1][col]
        && grid[row + 1].chars().nth(col).unwrap() == (current as u8 + 1) as char
    {
        count += traverse(grid, row + 1, col, visited);
    }

    if col > 0
        && !visited[row][col - 1]
        && grid[row].chars().nth(col - 1).unwrap() == (current as u8 + 1) as char
    {
        count += traverse(grid, row, col - 1, visited);
    }

    if col < cols - 1
        && !visited[row][col + 1]
        && grid[row].chars().nth(col + 1).unwrap() == (current as u8 + 1) as char
    {
        count += traverse(grid, row, col + 1, visited);
    }

    count
}

fn traverse_with_brute_force(
    grid: &Vec<String>,
    row: usize,
    col: usize,
    visited: &mut Vec<Vec<bool>>,
) -> i32 {
    let rows = grid.len();
    let cols = grid[0].len();

    let current = grid[row].chars().nth(col).unwrap();

    if current == '9' {
        return 1;
    }

    visited[row][col] = true;

    let mut count = 0;

    if row > 0
        && !visited[row - 1][col]
        && grid[row - 1].chars().nth(col).unwrap() == (current as u8 + 1) as char
    {
        count += traverse_with_brute_force(grid, row - 1, col, visited);
    }

    if row < rows - 1
        && !visited[row + 1][col]
        && grid[row + 1].chars().nth(col).unwrap() == (current as u8 + 1) as char
    {
        count += traverse_with_brute_force(grid, row + 1, col, visited);
    }

    if col > 0
        && !visited[row][col - 1]
        && grid[row].chars().nth(col - 1).unwrap() == (current as u8 + 1) as char
    {
        count += traverse_with_brute_force(grid, row, col - 1, visited);
    }

    if col < cols - 1
        && !visited[row][col + 1]
        && grid[row].chars().nth(col + 1).unwrap() == (current as u8 + 1) as char
    {
        count += traverse_with_brute_force(grid, row, col + 1, visited);
    }

    visited[row][col] = false;

    count
}

fn part_one(input: Vec<String>) -> io::Result<i32> {
    let rows = input.len();
    let cols = input[0].len();

    let mut result = 0;
    for row in 0..rows {
        for col in 0..cols {
            if input[row].chars().nth(col).unwrap() == '0' {
                let mut visited = vec![vec![false; cols]; rows];
                result += traverse(&input, row, col, &mut visited);
            }
        }
    }

    Ok(result)
}

fn part_two(input: Vec<String>) -> io::Result<i32> {
    let rows = input.len();
    let cols = input[0].len();

    let mut result = 0;
    for row in 0..rows {
        for col in 0..cols {
            if input[row].chars().nth(col).unwrap() == '0' {
                let mut visited = vec![vec![false; cols]; rows];
                result += traverse_with_brute_force(&input, row, col, &mut visited);
            }
        }
    }

    Ok(result)
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
