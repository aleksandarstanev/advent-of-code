use std::collections::{BTreeSet, HashMap};
use std::fs::File;
use std::io::{self, BufRead};

struct Input {
    grid: Vec<String>,
    start: (usize, usize),
    end: (usize, usize),
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Debug, Clone, Copy, Hash)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Debug, Clone, Copy, Hash)]
struct State {
    row: usize,
    col: usize,
    direction: Direction,
}

const TURN_COST: i32 = 1000;

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

fn dijkstra(grid: Vec<String>, start_states: Vec<State>) -> io::Result<HashMap<State, i32>> {
    let rows = grid.len();
    let cols = grid[0].len();

    let mut visited = vec![vec![vec![false; 4]; cols]; rows];
    let mut queue: BTreeSet<(i32, usize, usize, Direction)> = BTreeSet::new();

    let mut min_distances = HashMap::new();

    for state in start_states {
        queue.insert((0, state.row, state.col, state.direction));
    }

    while !queue.is_empty() {
        let (cost, row, col, direction) = queue.pop_first().unwrap();

        if visited[row][col][direction as usize] {
            continue;
        }

        visited[row][col][direction as usize] = true;

        min_distances.insert(
            State {
                row,
                col,
                direction,
            },
            cost,
        );

        match direction {
            Direction::Up => {
                if row > 0 && grid[row - 1].chars().nth(col).unwrap() != '#' {
                    queue.insert((cost + 1, row - 1, col, Direction::Up));
                }

                queue.insert((cost + TURN_COST, row, col, Direction::Left));
                queue.insert((cost + TURN_COST, row, col, Direction::Right));
            }

            Direction::Down => {
                if row < rows - 1 && grid[row + 1].chars().nth(col).unwrap() != '#' {
                    queue.insert((cost + 1, row + 1, col, Direction::Down));
                }

                queue.insert((cost + TURN_COST, row, col, Direction::Left));
                queue.insert((cost + TURN_COST, row, col, Direction::Right));
            }

            Direction::Left => {
                if col > 0 && grid[row].chars().nth(col - 1).unwrap() != '#' {
                    queue.insert((cost + 1, row, col - 1, Direction::Left));
                }

                queue.insert((cost + TURN_COST, row, col, Direction::Up));
                queue.insert((cost + TURN_COST, row, col, Direction::Down));
            }

            Direction::Right => {
                if col < cols - 1 && grid[row].chars().nth(col + 1).unwrap() != '#' {
                    queue.insert((cost + 1, row, col + 1, Direction::Right));
                }

                queue.insert((cost + TURN_COST, row, col, Direction::Up));
                queue.insert((cost + TURN_COST, row, col, Direction::Down));
            }
        }
    }

    Ok(min_distances)
}

fn get_opposite_direction(direction: Direction) -> Direction {
    match direction {
        Direction::Up => Direction::Down,
        Direction::Down => Direction::Up,
        Direction::Left => Direction::Right,
        Direction::Right => Direction::Left,
    }
}

fn part_one(input: &Input) -> io::Result<i32> {
    let min_distances = dijkstra(
        input.grid.clone(),
        vec![State {
            row: input.start.0,
            col: input.start.1,
            direction: Direction::Right,
        }],
    )?;

    let mut result = i32::MAX;

    for direction in [
        Direction::Up,
        Direction::Down,
        Direction::Left,
        Direction::Right,
    ] {
        let state = State {
            row: input.end.0,
            col: input.end.1,
            direction,
        };

        if let Some(cost) = min_distances.get(&state) {
            result = result.min(*cost);
        }
    }

    Ok(result)
}

fn part_two(input: &Input) -> io::Result<i32> {
    let min_distances_from_start = dijkstra(
        input.grid.clone(),
        vec![State {
            row: input.start.0,
            col: input.start.1,
            direction: Direction::Right,
        }],
    )?;

    let min_distances_from_end = dijkstra(
        input.grid.clone(),
        vec![
            State {
                row: input.end.0,
                col: input.end.1,
                direction: Direction::Left,
            },
            State {
                row: input.end.0,
                col: input.end.1,
                direction: Direction::Right,
            },
            State {
                row: input.end.0,
                col: input.end.1,
                direction: Direction::Up,
            },
            State {
                row: input.end.0,
                col: input.end.1,
                direction: Direction::Down,
            },
        ],
    )?;

    let mut min_cost = i32::MAX;

    for direction in [
        Direction::Up,
        Direction::Down,
        Direction::Left,
        Direction::Right,
    ] {
        min_cost = min_cost.min(
            *min_distances_from_start
                .get(&State {
                    row: input.end.0,
                    col: input.end.1,
                    direction,
                })
                .unwrap(),
        );
    }

    let rows = input.grid.len();
    let cols = input.grid[0].len();

    let mut result = 0;

    for row in 0..rows {
        for col in 0..cols {
            for direction in [
                Direction::Up,
                Direction::Down,
                Direction::Left,
                Direction::Right,
            ] {
                let cost_from_start = min_distances_from_start.get(&State {
                    row,
                    col,
                    direction,
                });

                if cost_from_start.is_none() {
                    continue;
                }

                let opposite_direction = get_opposite_direction(direction);

                let cost_from_end = min_distances_from_end.get(&State {
                    row,
                    col,
                    direction: opposite_direction,
                });

                if cost_from_end.is_none() {
                    continue;
                }

                let current_cost = *cost_from_start.unwrap() + *cost_from_end.unwrap();

                if current_cost == min_cost {
                    result += 1;

                    break;
                }
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

    // let result = part_one(&input)?;
    // println!("Part one: {}", result);

    let result = part_two(&input)?;
    println!("Part two: {}", result);

    Ok(())
}
