use std::collections::HashMap;
use std::{env,fs};

fn main() {
    let args: Vec<String> = env::args().collect();
    let contents = fs::read_to_string(&args[1]).unwrap();

    let (left, right): (Vec<i32>, Vec<i32>) = contents
        .lines()
        .map(|line| {
            let parts: Vec<&str> = line.split_whitespace().collect();
            (parts[0].parse::<i32>().unwrap(), parts[1].parse::<i32>().unwrap())
        })
        .unzip();

    let mut right_counts: HashMap<i32, i32> = HashMap::new();
    for num in right {
        *right_counts.entry(num).or_insert(0) += 1;
    }

    let similarity: i32 = left
        .iter()
        .map(|&num| num * right_counts.get(&num).unwrap_or(&0))
        .sum();

    println!("{:?}", similarity);
}
