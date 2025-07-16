use std::collections::VecDeque;
use std::fs::File;
use std::io::{self, BufRead};

const ROWS: usize = 71;
const COLUMNS: usize = 71;

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<(i32, i32)>> {
    let mut coordinates = Vec::new();

    for line in reader.lines() {
        let line = line?;
        let parts: Vec<&str> = line.split(',').collect();

        let x = parts[0].parse::<i32>().unwrap_or(0);
        let y = parts[1].parse::<i32>().unwrap_or(0);
        coordinates.push((x, y));
    }

    Ok(coordinates)
}

fn part_one(input: &Vec<(i32, i32)>, limit_bytes: i32) -> i32 {
    let mut grid = vec![vec!['.'; COLUMNS as usize]; ROWS as usize];
    let mut distance = vec![vec![-1; COLUMNS as usize]; ROWS as usize];

    for i in 0..limit_bytes {
        grid[input[i as usize].0 as usize][input[i as usize].1 as usize] = '#';
    }

    let mut queue = VecDeque::new();
    queue.push_back((0, 0));
    distance[0][0] = 0;

    while let Some((row, col)) = queue.pop_front() {
        let dist = distance[row][col];

        // println!("{} {}", row, col);

        if row == ROWS - 1 && col == COLUMNS - 1 {
            return dist;
        }

        if row > 0 && distance[row - 1][col] == -1 && grid[row - 1][col] == '.' {
            queue.push_back((row - 1, col));

            distance[row - 1][col] = dist + 1;
        }

        if row + 1 < ROWS && distance[row + 1][col] == -1 && grid[row + 1][col] == '.' {
            queue.push_back((row + 1, col));

            distance[row + 1][col] = dist + 1;
        }

        if col > 0 && distance[row][col - 1] == -1 && grid[row][col - 1] == '.' {
            queue.push_back((row, col - 1));

            distance[row][col - 1] = dist + 1;
        }

        if col + 1 < COLUMNS && distance[row][col + 1] == -1 && grid[row][col + 1] == '.' {
            queue.push_back((row, col + 1));

            distance[row][col + 1] = dist + 1;
        }
    }

    -1
}

fn part_two(input: Vec<(i32, i32)>) -> (i32, i32) {
    let mut left = 0;
    let mut right = input.len() - 1;

    let mut ans = (-1, -1);

    while left <= right {
        let mid = (left + right) / 2;

        if part_one(&input, mid as i32 + 1) != -1 {
            left = mid + 1;
        } else {
            ans = input[mid as usize];

            right = mid - 1;
        }
    }

    return ans;
}

fn main() -> io::Result<()> {
    let file_path = "inputz.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    // let result = part_one(&input, 1024);
    // println!("Part one: {}", result);

    let result = part_two(input);
    println!("Part two: {:?}", result);

    Ok(())
}
