use std::collections::{HashMap, HashSet};
use std::fs::File;
use std::io::{self, BufRead};

struct Input {
    grid: Vec<Vec<char>>,
    robot_position: (usize, usize),
    instructions: String,
}

fn parse_input(reader: io::BufReader<File>) -> io::Result<Input> {
    let mut grid = Vec::new();
    let mut robot_position = (0, 0);

    let mut lines = reader.lines();

    // Read grid
    while let Some(line) = lines.next() {
        let line = line?;
        if line.is_empty() {
            break;
        }

        for (j, c) in line.chars().enumerate() {
            if c == '@' {
                robot_position = (grid.len(), j);
            }
        }

        grid.push(line.chars().collect());
    }

    // Read instructions
    let mut instructions = String::new();
    while let Some(line) = lines.next() {
        let line = line?;
        instructions.push_str(&line);
    }

    Ok(Input {
        grid,
        robot_position,
        instructions,
    })
}

fn move_recursively(
    grid: &mut Vec<Vec<char>>,
    position: (usize, usize),
    row_delta: i32,
    col_delta: i32,
) -> (usize, usize) {
    let rows = grid.len();
    let cols = grid[0].len();

    let (row, col) = position;

    let new_row = row as i32 + row_delta;
    let new_col = col as i32 + col_delta;

    if new_row >= rows as i32
        || new_col >= cols as i32
        || new_row < 0
        || new_col < 0
        || grid[new_row as usize][new_col as usize] == '#'
    {
        return position;
    }

    let next_char = grid[new_row as usize][new_col as usize];
    if next_char == 'O' || next_char == '[' || next_char == ']' {
        move_recursively(
            grid,
            (new_row as usize, new_col as usize),
            row_delta,
            col_delta,
        );
    }

    if grid[new_row as usize][new_col as usize] == '.' {
        grid[new_row as usize][new_col as usize] = grid[row][col];
        grid[row][col] = '.';
        return (new_row as usize, new_col as usize);
    }

    position
}

fn move_robot(
    grid: &mut Vec<Vec<char>>,
    position: (usize, usize),
    direction: char,
) -> (usize, usize) {
    let (row, col) = position;

    let row_delta = match direction {
        'v' => 1,
        '^' => -1,
        _ => 0,
    };

    let col_delta = match direction {
        '>' => 1,
        '<' => -1,
        _ => 0,
    };

    move_recursively(grid, (row, col), row_delta, col_delta)
}

fn part_one(input: Input) -> i32 {
    let mut grid = input.grid.clone();
    let mut robot_position = input.robot_position;

    for direction in input.instructions.chars() {
        let new_position = move_robot(&mut grid, robot_position, direction);

        robot_position = new_position;
    }

    let mut result = 0;

    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            if grid[i][j] == 'O' {
                result += (100 * i + j) as i32;
            }
        }
    }

    result
}

fn enlarge_grid(grid: Vec<Vec<char>>) -> (Vec<Vec<char>>, (usize, usize)) {
    let mut new_grid = Vec::new();

    let mut robot_position = (0, 0);
    for row in grid {
        let mut new_row = Vec::new();

        for c in row {
            match c {
                '#' => {
                    new_row.push('#');
                    new_row.push('#');
                }
                '.' => {
                    new_row.push('.');
                    new_row.push('.');
                }
                'O' => {
                    new_row.push('[');
                    new_row.push(']');
                }
                '@' => {
                    robot_position = (new_grid.len(), new_row.len());
                    new_row.push('@');
                    new_row.push('.');
                }
                _ => unreachable!(),
            }
        }

        new_grid.push(new_row);
    }

    (new_grid, robot_position)
}

fn can_move_in_enlarged_grid(
    grid: &mut Vec<Vec<char>>,
    position: (usize, usize),
    row_delta: i32,
    memo: &mut HashMap<(usize, usize), bool>,
) -> bool {
    let rows = grid.len();

    let (row, col) = position;

    if memo.contains_key(&position) {
        return memo[&position];
    }

    let new_row = row as i32 + row_delta;

    if new_row >= rows as i32 || new_row < 0 || grid[new_row as usize][col] == '#' {
        memo.insert(position, false);
        return false;
    }

    if grid[new_row as usize][col] == '.' {
        memo.insert(position, true);
        return true;
    }

    if grid[new_row as usize][col] == ']' {
        let result = can_move_in_enlarged_grid(grid, (new_row as usize, col), row_delta, memo)
            && can_move_in_enlarged_grid(grid, (new_row as usize, col - 1), row_delta, memo);
        memo.insert(position, result);

        return result;
    }

    if grid[new_row as usize][col] == '[' {
        let result = can_move_in_enlarged_grid(grid, (new_row as usize, col), row_delta, memo)
            && can_move_in_enlarged_grid(grid, (new_row as usize, col + 1), row_delta, memo);

        memo.insert(position, result);
        return result;
    }

    memo.insert(position, false);
    false
}

fn move_recursively_in_enlarged_grid(
    grid: &mut Vec<Vec<char>>,
    position: (usize, usize),
    row_delta: i32,
    memo: &mut HashSet<(usize, usize)>,
) -> (usize, usize) {
    let (row, col) = position;

    if memo.contains(&position) {
        return position;
    }

    memo.insert(position);

    let new_row = (row as i32 + row_delta) as usize;

    if grid[new_row as usize][col] == '[' {
        move_recursively_in_enlarged_grid(grid, (new_row as usize, col), row_delta, memo);
        move_recursively_in_enlarged_grid(grid, (new_row as usize, col + 1), row_delta, memo);
    }

    if grid[new_row as usize][col] == ']' {
        move_recursively_in_enlarged_grid(grid, (new_row as usize, col), row_delta, memo);
        move_recursively_in_enlarged_grid(grid, (new_row as usize, col - 1), row_delta, memo);
    }

    if grid[new_row as usize][col] == '.' {
        grid[new_row as usize][col] = grid[row][col];
        grid[row][col] = '.';
        return (new_row as usize, col);
    }

    position
}

fn move_robot_in_enlarged_grid(
    grid: &mut Vec<Vec<char>>,
    position: (usize, usize),
    direction: char,
) -> (usize, usize) {
    let (row, col) = position;

    let row_delta = match direction {
        'v' => 1,
        '^' => -1,
        _ => 0,
    };

    let mut memo = HashMap::new();

    if can_move_in_enlarged_grid(grid, (row, col), row_delta, &mut memo) {
        let mut memo = HashSet::new();

        return move_recursively_in_enlarged_grid(grid, (row, col), row_delta, &mut memo);
    }

    position
}

fn part_two(input: Input) -> i32 {
    let (mut grid, mut robot_position) = enlarge_grid(input.grid.clone());

    for direction in input.instructions.chars() {
        match direction {
            '>' | '<' => {
                robot_position = move_robot(&mut grid, robot_position, direction);
            }
            'v' | '^' => {
                robot_position = move_robot_in_enlarged_grid(&mut grid, robot_position, direction);
            }
            _ => unreachable!(),
        }

        // for row in &grid {
        //     println!("{}", row.iter().collect::<String>());
        // }

        // println!("");
    }

    let mut result = 0;

    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            if grid[i][j] == '[' {
                result += (100 * i + j) as i32;
            }
        }
    }

    result
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    // let result = part_one(input);
    // println!("{}", result);

    let result = part_two(input);
    println!("{}", result);

    Ok(())
}
