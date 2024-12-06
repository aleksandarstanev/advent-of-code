use std::collections::HashSet;
use std::fs::File;
use std::io::{self, BufRead};

struct Rule {
    first: String,
    second: String,
}

struct Input {
    rules: Vec<Rule>,

    updates: Vec<Vec<String>>,
}

fn parse_input(reader: io::BufReader<File>) -> io::Result<Input> {
    let mut rules: Vec<Rule> = Vec::new();

    let lines: Vec<String> = reader.lines().collect::<Result<_, _>>()?;

    let mut lines_iter = lines.iter();

    for line in &mut lines_iter {
        if line == "" {
            break;
        }

        let parts: Vec<&str> = line.split_terminator('|').collect();

        rules.push(Rule {
            first: parts[0].to_string(),
            second: parts[1].to_string(),
        });
    }

    let mut updates: Vec<Vec<String>> = Vec::new();

    for line in lines_iter {
        let pages: Vec<String> = line.split_terminator(',').map(|s| s.to_string()).collect();

        updates.push(pages);
    }

    return Ok(Input { rules, updates });
}

fn part_one(input: Input) -> io::Result<i32> {
    let mut result = 0;

    for update in &input.updates {
        let mut is_valid = true;

        for rule in &input.rules {
            let first = &rule.first;
            let second = &rule.second;

            let mut first_idx: usize = usize::MAX;
            let mut second_idx: usize = usize::MAX;

            for (idx, page) in update.iter().enumerate() {
                if page == first {
                    first_idx = idx;
                }

                if page == second {
                    second_idx = idx;
                }
            }

            if first_idx != usize::MAX && second_idx != usize::MAX && first_idx > second_idx {
                is_valid = false;
                break;
            }
        }

        if is_valid {
            result += update[update.len() / 2].to_string().parse::<i32>().unwrap();
        }
    }

    return Ok(result);
}

fn visit(
    node: String,
    visited: &mut HashSet<String>,
    input: &Input,
    top_sort: &mut Vec<String>,
    update: &Vec<String>,
) {
    if visited.contains(&node) {
        return;
    }

    visited.insert(node.clone());

    for rule in &input.rules {
        if rule.first == node && update.contains(&rule.second) {
            visit(rule.second.clone(), visited, input, top_sort, update);
        }
    }

    top_sort.push(node);
}

fn part_two(input: Input) -> io::Result<i32> {
    let mut result = 0;

    for update in &input.updates {
        let mut is_valid = true;

        for rule in &input.rules {
            let first = &rule.first;
            let second = &rule.second;

            let mut first_idx: usize = usize::MAX;
            let mut second_idx: usize = usize::MAX;

            for (idx, page) in update.iter().enumerate() {
                if page == first {
                    first_idx = idx;
                }

                if page == second {
                    second_idx = idx;
                }
            }

            if first_idx != usize::MAX && second_idx != usize::MAX && first_idx > second_idx {
                is_valid = false;
                break;
            }
        }

        if is_valid {
            continue;
        }

        let mut top_sort = Vec::new();
        let mut visited: HashSet<String> = HashSet::new();
        for i in 0..update.len() {
            if !visited.contains(&update[i]) {
                visit(
                    update[i].clone(),
                    &mut visited,
                    &input,
                    &mut top_sort,
                    &update,
                );
            }
        }

        result += top_sort[update.len() / 2]
            .to_string()
            .parse::<i32>()
            .unwrap();
    }

    return Ok(result);
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

// 4667 -> too low
