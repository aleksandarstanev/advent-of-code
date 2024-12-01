use std::{
    fs::File,
    io::{self, BufRead},
};

fn main() -> io::Result<()> {
    let file_path = "input.txt";
    let file = File::open(&file_path)?;

    let reader = io::BufReader::new(file);

    let mut result = 0;
    for line in reader.lines() {
        match line {
            Ok(line_content) => {
                let mut first_digit: i32 = -1;
                let mut last_digit: i32 = -1;

                for (_, character) in line_content.chars().enumerate() {
                    if character.is_digit(10) {
                        if first_digit == -1 {
                            match character.to_digit(10) {
                                Some(digit) => first_digit = digit as i32,
                                None => eprintln!("Error: {} is not a digit", character),
                            }
                        }

                        match character.to_digit(10) {
                            Some(digit) => last_digit = digit as i32,
                            None => eprintln!("Error: {} is not a digit", character),
                        }
                    }
                }

                let current_number = first_digit * 10 + last_digit;

                result = result + current_number;
            }
            Err(err) => eprintln!("Error: {}", err),
        }
    }

    println!("Result: {}", result);

    Ok(())
}
