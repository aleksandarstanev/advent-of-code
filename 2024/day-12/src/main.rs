use std::fs::File;
use std::io::{self, BufRead};

fn parse_input(reader: io::BufReader<File>) -> io::Result<Vec<String>> {
    let mut lines = Vec::new();
    for line in reader.lines() {
        lines.push(line?);
    }

    Ok(lines)
}

fn get_area_and_perimeter(
    grid: &Vec<String>,
    row: usize,
    col: usize,
    visited: &mut Vec<Vec<bool>>,
) -> io::Result<(i32, i32)> {
    let rows = grid.len();
    let cols = grid[0].len();

    if visited[row][col] {
        return Ok((0, 0));
    }

    let current_cell = grid[row].chars().nth(col).unwrap();

    let mut area = 1;
    let mut perimeter = 0;

    visited[row][col] = true;

    if row > 0 && grid[row - 1].chars().nth(col).unwrap() == current_cell {
        let (a, p) = get_area_and_perimeter(grid, row - 1, col, visited)?;
        area += a;
        perimeter += p;
    } else {
        perimeter += 1;
    }

    if row < rows - 1 && grid[row + 1].chars().nth(col).unwrap() == current_cell {
        let (a, p) = get_area_and_perimeter(grid, row + 1, col, visited)?;
        area += a;
        perimeter += p;
    } else {
        perimeter += 1;
    }

    if col > 0 && grid[row].chars().nth(col - 1).unwrap() == current_cell {
        let (a, p) = get_area_and_perimeter(grid, row, col - 1, visited)?;
        area += a;
        perimeter += p;
    } else {
        perimeter += 1;
    }

    if col < cols - 1 && grid[row].chars().nth(col + 1).unwrap() == current_cell {
        let (a, p) = get_area_and_perimeter(grid, row, col + 1, visited)?;
        area += a;
        perimeter += p;
    } else {
        perimeter += 1;
    }

    Ok((area, perimeter))
}

fn get_sides_and_area(
    grid: &Vec<String>,
    row: usize,
    col: usize,
    visited: &mut Vec<Vec<bool>>,
) -> io::Result<(i32, i32)> {
    let rows = grid.len();
    let cols = grid[0].len();

    if visited[row][col] {
        return Ok((0, 0));
    }

    let current_cell = grid[row].chars().nth(col).unwrap();

    let mut area = 1;
    let mut sides = 0;

    visited[row][col] = true;

    if row > 0 && grid[row - 1].chars().nth(col).unwrap() == current_cell {
        let (a, s) = get_sides_and_area(grid, row - 1, col, visited)?;
        area += a;
        sides += s;
    }

    if row < rows - 1 && grid[row + 1].chars().nth(col).unwrap() == current_cell {
        let (a, s) = get_sides_and_area(grid, row + 1, col, visited)?;
        area += a;
        sides += s;
    }

    if col > 0 && grid[row].chars().nth(col - 1).unwrap() == current_cell {
        let (a, s) = get_sides_and_area(grid, row, col - 1, visited)?;
        area += a;
        sides += s;
    }

    if col < cols - 1 && grid[row].chars().nth(col + 1).unwrap() == current_cell {
        let (a, s) = get_sides_and_area(grid, row, col + 1, visited)?;
        area += a;
        sides += s;
    }

    if row == 0 || grid[row - 1].chars().nth(col).unwrap() != current_cell {
        let is_left_same = col > 0
            && (grid[row].chars().nth(col - 1).unwrap() == current_cell
                && (row == 0 || grid[row - 1].chars().nth(col - 1).unwrap() != current_cell));

        if !is_left_same {
            sides += 1;
        }
    }

    if row == rows - 1 || grid[row + 1].chars().nth(col).unwrap() != current_cell {
        let is_left_same = col > 0
            && (grid[row].chars().nth(col - 1).unwrap() == current_cell
                && (row == rows - 1
                    || grid[row + 1].chars().nth(col - 1).unwrap() != current_cell));

        if !is_left_same {
            sides += 1;
        }
    }

    if col == 0 || grid[row].chars().nth(col - 1).unwrap() != current_cell {
        let is_top_same = row > 0
            && (grid[row - 1].chars().nth(col).unwrap() == current_cell
                && (col == 0 || grid[row - 1].chars().nth(col - 1).unwrap() != current_cell));

        if !is_top_same {
            sides += 1;
        }
    }

    if col == cols - 1 || grid[row].chars().nth(col + 1).unwrap() != current_cell {
        let is_bottom_same = row < rows - 1
            && (grid[row + 1].chars().nth(col).unwrap() == current_cell
                && (col == cols - 1
                    || grid[row + 1].chars().nth(col + 1).unwrap() != current_cell));

        if !is_bottom_same {
            sides += 1;
        }
    }

    Ok((area, sides))
}

fn part_one(grid: Vec<String>) -> io::Result<i32> {
    let rows = grid.len();
    let cols = grid[0].len();

    let mut visited = vec![vec![false; cols]; rows];

    let mut result = 0;

    for row in 0..rows {
        for col in 0..cols {
            if !visited[row][col] {
                let (area, perimeter) = get_area_and_perimeter(&grid, row, col, &mut visited)?;
                result += area * perimeter;
            }
        }
    }

    Ok(result)
}

fn part_two(grid: Vec<String>) -> io::Result<i32> {
    let rows = grid.len();
    let cols = grid[0].len();

    let mut visited = vec![vec![false; cols]; rows];

    let mut result = 0;

    for row in 0..rows {
        for col in 0..cols {
            if !visited[row][col] {
                let (sides, perimeter) = get_sides_and_area(&grid, row, col, &mut visited)?;
                result += sides * perimeter;
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
