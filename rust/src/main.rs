use std::fs::{read_to_string, File};
use std::io::{self, BufRead, Write};
use std::process::exit;

const OUTPUT_FILE: &str = "output.txt";
const NEIGHBORS: [(i32, i32); 8] = [
    (-1, -1),
    (-1, 0),
    (-1, 1),
    (0, -1),
    (0, 1),
    (1, -1),
    (1, 0),
    (1, 1),
];

fn parse_file(filename: &str) -> Vec<Vec<i8>> {
    let mut result = Vec::new();
    for line in read_to_string(filename).unwrap().lines() {
        let row: Vec<i8> = line
            .chars()
            .filter_map(|c| c.to_digit(10).map(|d| d as i8))
            .collect();
        result.push(row);
    }
    result
}

fn next_state(board: &Vec<Vec<i8>>) -> Vec<Vec<i8>> {
    let mut new_board = board.clone();
    for y in 0..board.len() {
        for x in 0..board[0].len() {
            new_board[y][x] = next_cell(board, y as i32, x as i32);
        }
    }
    new_board
}

fn next_cell(board: &Vec<Vec<i8>>, y: i32, x: i32) -> i8 {
    let mut live_neighbors = 0;
    let height = board.len() as i32;
    let width = board[0].len() as i32;

    for (dy, dx) in NEIGHBORS.iter() {
        let ny = (y + dy + height) % height;
        let nx = (x + dx + width) % width;
        live_neighbors += board[ny as usize][nx as usize];
    }

    match board[y as usize][x as usize] {
        1 if live_neighbors < 2 || live_neighbors > 3 => 0,
        0 if live_neighbors == 3 => 1,
        state => state,
    }
}

fn print_board(board: &Vec<Vec<i8>>) {
    let mut file = File::create(OUTPUT_FILE).expect("Unable to create file");
    let mut buffer = String::new();
    for row in board {
        for &cell in row {
            buffer.push(if cell == 1 { '1' } else { '0' });
        }
        buffer.push('\n');
    }
    file.write_all(buffer.as_bytes())
        .expect("Unable to write data");
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    let mut input_file = None;
    let mut amount: Option<i32> = None;
    let mut silent = false;
    let mut server_mode = false;

    let mut args_iter = args.iter().skip(1);
    while let Some(arg) = args_iter.next() {
        match arg.as_str() {
            "--file" => input_file = args_iter.next(),
            "--amount" => amount = args_iter.next().and_then(|s| s.parse().ok()),
            "--silent" => silent = true,
            "--server" => server_mode = true,
            _ => {
                eprintln!("Unknown argument: {}", arg);
                exit(1);
            }
        }
    }

    if server_mode {
        println!("READY");
        let stdin = io::stdin();
        for line in stdin.lock().lines() {
            match line {
                Ok(l) => {
                    let l = l.trim();
                    if l.is_empty() {
                        continue;
                    }
                    if l == "SHUTDOWN" {
                        break;
                    }
                    if l.starts_with("RUN ") {
                        let rest = &l[4..];
                        let mut parts = rest.split_whitespace();
                        if let Some(f) = parts.next() {
                            let n = parts
                                .next()
                                .and_then(|s| s.parse::<i32>().ok())
                                .unwrap_or(1);
                            let mut b = parse_file(f);
                            for _ in 0..n {
                                b = next_state(&b);
                            }
                            if !silent {
                                print_board(&b);
                            }
                            println!("DONE");
                        } else {
                            println!("ERROR bad request");
                        }
                    } else {
                        println!("ERROR unknown command");
                    }
                }
                Err(_) => break,
            }
        }
    } else {
        let input_file = match input_file {
            Some(file) => file,
            None => {
                eprintln!("No input file specified");
                exit(1);
            }
        };

        if let Some(amount) = amount {
            let mut current_board = parse_file(&input_file);
            for _ in 0..amount {
                current_board = next_state(&current_board);
            }
            if !silent {
                print_board(&current_board);
            }
        }
    }
}
