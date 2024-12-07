use std::collections::HashSet;
use std::fs::File;
use std::io::{self, BufRead};

#[derive(Copy, Clone, Eq, Hash, PartialEq)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<String>> {
    reader.lines().collect()
}

fn get_current_position(input: &Vec<String>) -> (i32, i32) {
    let rows = input.len();
    let cols = input[0].len();

    for i in 0..rows {
        for j in 0..cols {
            if input[i].chars().nth(j).unwrap() == '^' {
                return (i as i32, j as i32);
            }
        }
    }

    return (0, 0);
}

fn part_one(input: Vec<String>) -> io::Result<i32> {
    let rows = input.len();
    let cols = input[0].len();

    let (mut cur_row, mut cur_col) = get_current_position(&input);

    let mut steps = 0;
    let mut direction = Direction::Up;
    let mut visited: HashSet<(i32, i32)> = HashSet::new();

    while cur_row >= 0 && cur_col >= 0 && cur_row < rows as i32 && cur_col < cols as i32 {
        if input[cur_row as usize]
            .chars()
            .nth(cur_col as usize)
            .unwrap()
            == '#'
        {
            match direction {
                Direction::Up => {
                    cur_row += 1;
                    direction = Direction::Right;
                }
                Direction::Down => {
                    cur_row -= 1;
                    direction = Direction::Left;
                }
                Direction::Left => {
                    cur_col += 1;
                    direction = Direction::Up;
                }
                Direction::Right => {
                    cur_col -= 1;
                    direction = Direction::Down;
                }
            }

            continue;
        }

        if !visited.contains(&(cur_row, cur_col)) {
            steps += 1;
            visited.insert((cur_row, cur_col));
        }

        match direction {
            Direction::Up => {
                cur_row -= 1;
            }
            Direction::Down => {
                cur_row += 1;
            }
            Direction::Left => {
                cur_col -= 1;
            }
            Direction::Right => {
                cur_col += 1;
            }
        }
    }

    return Ok(steps);
}

fn has_cycle(input: Vec<String>) -> bool {
    let rows = input.len();
    let cols = input[0].len();

    let (mut cur_row, mut cur_col) = get_current_position(&input);

    let mut direction = Direction::Up;
    let mut visited: HashSet<(i32, i32, Direction)> = HashSet::new();

    while cur_row >= 0 && cur_col >= 0 && cur_row < rows as i32 && cur_col < cols as i32 {
        if input[cur_row as usize]
            .chars()
            .nth(cur_col as usize)
            .unwrap()
            == '#'
        {
            match direction {
                Direction::Up => {
                    cur_row += 1;
                    direction = Direction::Right;
                }
                Direction::Down => {
                    cur_row -= 1;
                    direction = Direction::Left;
                }
                Direction::Left => {
                    cur_col += 1;
                    direction = Direction::Up;
                }
                Direction::Right => {
                    cur_col -= 1;
                    direction = Direction::Down;
                }
            }

            continue;
        }

        if visited.contains(&(cur_row, cur_col, direction)) {
            return true;
        }

        visited.insert((cur_row, cur_col, direction));

        match direction {
            Direction::Up => {
                cur_row -= 1;
            }
            Direction::Down => {
                cur_row += 1;
            }
            Direction::Left => {
                cur_col -= 1;
            }
            Direction::Right => {
                cur_col += 1;
            }
        }
    }

    return false;
}

fn part_two(input: Vec<String>) -> io::Result<i32> {
    let rows = input.len();
    let cols = input[0].len();

    let mut cycles = 0;

    for i in 0..rows {
        for j in 0..cols {
            if input[i].chars().nth(j).unwrap() == '.' {
                let mut input_cloned = input.clone();
                let mut chars: Vec<char> = input_cloned[i].chars().collect();
                chars[j] = '#';
                input_cloned[i] = chars.into_iter().collect();

                if has_cycle(input_cloned) {
                    cycles += 1;
                }
            }
        }
    }

    return Ok(cycles);
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    // let result = part_one(input).unwrap();
    // println!("Part One: {}", result);

    let result = part_two(input).unwrap();
    println!("Part Two: {}", result);

    Ok(())
}
