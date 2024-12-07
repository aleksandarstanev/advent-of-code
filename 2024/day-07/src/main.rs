use std::fs::File;
use std::io::{self, BufRead};

#[derive(Debug)]
struct Equation {
    result: i64,
    numbers: Vec<i64>,
}

#[derive(Debug)]
struct Input {
    equations: Vec<Equation>,
}

fn parse_input(reader: io::BufReader<File>) -> io::Result<Input> {
    let mut equations = Vec::new();
    for line in reader.lines() {
        let line = line?;
        let parts: Vec<&str> = line.split_terminator(':').collect();

        let result = parts[0].parse::<i64>().unwrap();
        let numbers = parts[1]
            .split_whitespace()
            .map(|x| x.trim().parse::<i64>().unwrap())
            .collect();

        equations.push(Equation { result, numbers });
    }

    Ok(Input { equations })
}

fn can_solve(equation: &Equation, acc: i64, target: i64, idx: usize, with_concat: bool) -> bool {
    if idx == equation.numbers.len() {
        return acc == target;
    }

    let next_sum = if idx == 0 {
        equation.numbers[0]
    } else {
        acc + equation.numbers[idx]
    };

    let next_mul = if idx == 0 {
        equation.numbers[0]
    } else {
        acc * equation.numbers[idx]
    };

    let mut res = can_solve(equation, next_sum, target, idx + 1, with_concat)
        || can_solve(equation, next_mul, target, idx + 1, with_concat);

    if with_concat {
        let next_concat = if idx == 0 {
            equation.numbers[0]
        } else {
            (acc.to_string() + equation.numbers[idx].to_string().as_str())
                .parse()
                .unwrap()
        };

        res = res || can_solve(equation, next_concat, target, idx + 1, with_concat)
    }

    return res;
}

fn part_one(input: Input) -> io::Result<i64> {
    let mut total: i64 = 0;
    for equation in input.equations {
        if can_solve(&equation, 0, equation.result, 0, false) {
            total += equation.result;
        }
    }

    Ok(total)
}

fn part_two(input: Input) -> io::Result<i64> {
    let mut total: i64 = 0;
    for equation in input.equations {
        if can_solve(&equation, 0, equation.result, 0, true) {
            total += equation.result;
        }
    }

    Ok(total)
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
