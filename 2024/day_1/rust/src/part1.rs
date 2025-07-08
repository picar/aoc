use std::{env,fs, vec};

fn main() {
    let args: Vec<String> = env::args().collect();
    let contents = fs::read_to_string(&args[1]).unwrap();

    let mut left: Vec<i32> = vec![];
    let mut right: Vec<i32>  = vec![];

    for l in contents.lines() {
        let parts: Vec<&str> = l.split_whitespace().collect();
        left.push(parts[0].parse::<i32>().unwrap());
        right.push(parts[1].parse::<i32>().unwrap());
    }
    left.sort();
    right.sort();

    let s: i32 = left.iter()
        .zip(right.iter())
        .map(|(l, r)| (r - l).abs())
        .sum();
    println!("{:?}", s);
}