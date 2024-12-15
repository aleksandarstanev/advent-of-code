use std::fs::File;
use std::io::{self, BufRead};

use scan_fmt::scan_fmt;

#[derive(Debug)]
struct Robot {
    start_x: i32,
    start_y: i32,

    x_delta: i32,
    y_delta: i32,
}

#[derive(Debug)]
struct Input {
    robots: Vec<Robot>,
}

fn parse_input(reader: io::BufReader<File>) -> io::Result<Input> {
    let mut robots = Vec::new();
    let mut lines = reader.lines();

    while let Some(line) = lines.next() {
        let line = line?;
        let (start_x, start_y, x_delta, y_delta) =
            scan_fmt!(&line, "p={},{} v={},{}", i32, i32, i32, i32).unwrap();

        robots.push(Robot {
            start_x,
            start_y,
            x_delta,
            y_delta,
        });
    }

    Ok(Input { robots })
}

fn part_one(input: Input) -> i32 {
    let total_rows = 103;
    let total_cols = 101;
    let iterations = 100;

    let mut quadrant_top_left = 0;
    let mut quadrant_top_right = 0;
    let mut quadrant_bottom_left = 0;
    let mut quadrant_bottom_right = 0;

    for robot in input.robots {
        let final_x =
            ((robot.start_x + robot.x_delta * iterations) % total_cols + total_cols) % total_cols;
        let final_y =
            ((robot.start_y + robot.y_delta * iterations) % total_rows + total_rows) % total_rows;

        if final_x < total_cols / 2 && final_y < total_rows / 2 {
            quadrant_top_left += 1;
        } else if final_x < total_cols / 2 && final_y > total_rows / 2 {
            quadrant_bottom_left += 1;
        } else if final_x > total_cols / 2 && final_y < total_rows / 2 {
            quadrant_top_right += 1;
        } else if final_x > total_cols / 2 && final_y > total_rows / 2 {
            quadrant_bottom_right += 1;
        }
    }

    quadrant_top_left * quadrant_bottom_right * quadrant_top_right * quadrant_bottom_left
}

fn display_grid(grid: &Vec<Vec<bool>>) {
    for row in grid {
        for cell in row {
            print!("{}", if *cell { '#' } else { '.' });
        }
        println!();
    }
}

fn check_adjacent_robots(grid: &Vec<Vec<bool>>) -> i32 {
    let mut count = 0;

    for (r, row) in grid.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            if *cell {
                let has_adjacent_robot = r > 0 && grid[r - 1][c]
                    || r < grid.len() - 1 && grid[r + 1][c]
                    || c > 0 && grid[r][c - 1]
                    || c < grid[0].len() - 1 && grid[r][c + 1]
                    || r > 0 && c > 0 && grid[r - 1][c - 1]
                    || r > 0 && c < grid[0].len() - 1 && grid[r - 1][c + 1]
                    || r < grid.len() - 1 && c > 0 && grid[r + 1][c - 1]
                    || r < grid.len() - 1 && c < grid[0].len() - 1 && grid[r + 1][c + 1];

                if has_adjacent_robot {
                    count += 1;
                }
            }
        }
    }

    count
}

fn part_two(input: Input) {
    let max_iterations = 100000;
    let total_rows = 103;
    let total_cols = 101;

    for iteration in 1..max_iterations {
        let mut grid = vec![vec![false; 101]; 103];

        for robot in &input.robots {
            let final_x = ((robot.start_x + robot.x_delta * iteration) % total_cols + total_cols)
                % total_cols;
            let final_y = ((robot.start_y + robot.y_delta * iteration) % total_rows + total_rows)
                % total_rows;

            grid[final_y as usize][final_x as usize] = true;
        }

        let robots_with_adjacent = check_adjacent_robots(&grid);

        if robots_with_adjacent > 300 {
            println!("{}", robots_with_adjacent);
            println!("{}", iteration);
            display_grid(&grid);
        }
    }
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    // let result = part_one(input);

    // println!("{}", result);

    part_two(input);

    Ok(())
}
