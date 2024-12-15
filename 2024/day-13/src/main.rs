use scan_fmt::scan_fmt;
use std::fs::File;
use std::io::{self, BufRead};

#[derive(Debug)]
struct Machine {
    x_diff_a: i32,
    y_diff_a: i32,
    x_diff_b: i32,
    y_diff_b: i32,
    prize_x: i32,
    prize_y: i32,
}

#[derive(Debug)]
struct Input {
    machines: Vec<Machine>,
}

fn parse_input(reader: io::BufReader<File>) -> io::Result<Input> {
    let mut machines = Vec::new();
    let mut lines = reader.lines();

    while let Some(line_a) = lines.next() {
        let line_a = line_a?;
        let (x_diff_a, y_diff_a) = scan_fmt!(&line_a, "Button A: X+{}, Y+{}", i32, i32).unwrap();

        let line_b = lines.next().unwrap()?;
        let (x_diff_b, y_diff_b) = scan_fmt!(&line_b, "Button B: X+{}, Y+{}", i32, i32).unwrap();

        let line_prize = lines.next().unwrap()?;
        let (prize_x, prize_y) = scan_fmt!(&line_prize, "Prize: X={}, Y={}", i32, i32).unwrap();

        // Skip empty line if it exists
        lines.next();

        machines.push(Machine {
            x_diff_a,
            y_diff_a,
            x_diff_b,
            y_diff_b,
            prize_x,
            prize_y,
        });
    }

    Ok(Input { machines })
}

fn solve(input: Input, prize_diff_x: i64, prize_diff_y: i64, max_t1: i64, max_t2: i64) -> i64 {
    let mut result: i64 = 0;

    for machine in input.machines {
        let x_diff_a = machine.x_diff_a as i64;
        let y_diff_a = machine.y_diff_a as i64;
        let x_diff_b = machine.x_diff_b as i64;
        let y_diff_b = machine.y_diff_b as i64;
        let prize_x = machine.prize_x as i64 + prize_diff_x;
        let prize_y = machine.prize_y as i64 + prize_diff_y;

        let num = prize_y * x_diff_a - prize_x * y_diff_a;
        let denom = x_diff_a * y_diff_b - x_diff_b * y_diff_a;

        if num % denom != 0 {
            continue;
        }

        let t2 = num / denom;

        if (prize_x - x_diff_b * t2) % x_diff_a != 0 {
            continue;
        }

        let t1 = (prize_x - x_diff_b * t2) / x_diff_a;

        if t1 > max_t1 || t2 > max_t2 {
            continue;
        }

        result += t1 * 3 + t2;
    }

    result
}

fn part_one(input: Input) -> i64 {
    solve(input, 0, 0, 100, 100)
}

fn part_two(input: Input) -> i64 {
    solve(input, 10000000000000, 10000000000000, i64::MAX, i64::MAX)
}

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let input = parse_input(reader)?;

    let result = part_two(input);

    println!("{}", result);

    Ok(())
}
