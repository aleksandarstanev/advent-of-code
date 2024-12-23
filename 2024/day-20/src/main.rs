use std::collections::VecDeque;
use std::fs::File;
use std::io::{self, BufRead};

struct Input {
    grid: Vec<String>,
    start: (usize, usize),
    end: (usize, usize),
}

fn parse_input(reader: io::BufReader<File>) -> io::Result<Input> {
    let mut lines = Vec::new();
    for line in reader.lines() {
        lines.push(line?);
    }

    let mut start = (0, 0);
    let mut end = (0, 0);

    for (i, row) in lines.iter().enumerate() {
        for (j, col) in row.chars().enumerate() {
            if col == 'S' {
                start = (i, j);
            }

            if col == 'E' {
                end = (i, j);
            }
        }
    }

    Ok(Input {
        grid: lines,
        start,
        end,
    })
}

fn bfs(grid: &Vec<String>, start: (usize, usize)) -> Vec<Vec<i32>> {
    let rows = grid.len();
    let cols = grid[0].len();

    let grid = &grid;

    let mut distances = vec![vec![-1; rows]; cols];

    let mut queue = VecDeque::new();
    queue.push_back((start.0, start.1));
    distances[start.0][start.1] = 0;

    while let Some((row, col)) = queue.pop_front() {
        let dist = distances[row][col];

        if row > 0
            && grid[row - 1].chars().nth(col).unwrap() != '#'
            && distances[row - 1][col] == -1
        {
            queue.push_back((row - 1, col));
            distances[row - 1][col] = dist + 1;
        }

        if row + 1 < rows
            && grid[row + 1].chars().nth(col).unwrap() != '#'
            && distances[row + 1][col] == -1
        {
            queue.push_back((row + 1, col));
            distances[row + 1][col] = dist + 1;
        }

        if col > 0
            && grid[row].chars().nth(col - 1).unwrap() != '#'
            && distances[row][col - 1] == -1
        {
            queue.push_back((row, col - 1));
            distances[row][col - 1] = dist + 1;
        }

        if col + 1 < cols
            && grid[row].chars().nth(col + 1).unwrap() != '#'
            && distances[row][col + 1] == -1
        {
            queue.push_back((row, col + 1));
            distances[row][col + 1] = dist + 1;
        }
    }

    distances
}

fn solve(input: Input, cheats_max_dist: i32) -> i32 {
    let rows = input.grid.len();
    let cols = input.grid[0].len();

    let dist_from_start = bfs(&input.grid, input.start);
    let dist_from_end = bfs(&input.grid, input.end);

    let original_distance = dist_from_start[input.end.0][input.end.1];
    let mut result = 0;

    for i in 0..rows {
        for j in 0..cols {
            if input.grid[i].chars().nth(j).unwrap() == '#' {
                continue;
            }

            for ii in -cheats_max_dist..(cheats_max_dist + 1) {
                for jj in -cheats_max_dist..(cheats_max_dist + 1) {
                    if ii.abs() + jj.abs() > cheats_max_dist {
                        continue;
                    }

                    let new_i = i as i32 + ii;
                    let new_j = j as i32 + jj;

                    if new_i < 0 || new_i >= rows as i32 || new_j < 0 || new_j >= cols as i32 {
                        continue;
                    }

                    if input.grid[new_i as usize]
                        .chars()
                        .nth(new_j as usize)
                        .unwrap()
                        == '#'
                    {
                        continue;
                    }

                    if dist_from_start[i][j] != -1
                        && dist_from_end[new_i as usize][new_j as usize] != -1
                    {
                        let new_dist = dist_from_start[i][j]
                            + dist_from_end[new_i as usize][new_j as usize]
                            + ii.abs()
                            + jj.abs();

                        if original_distance - new_dist >= 100 {
                            result += 1;
                        }
                    }
                }
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

    // let result = solve(input, 2);

    // println!("Part one: {}", result);

    let result = solve(input, 20);

    println!("Part two: {}", result);

    Ok(())
}
