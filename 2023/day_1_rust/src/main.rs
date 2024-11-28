use std::fs::File;                                                            
use std::io::{self, BufRead};                                                 
use std::path::Path;                                                          
    

fn first_number(s: &str) -> Result<char, String> {
    for c in s.chars() {
        if c.is_digit(10) {
            return Ok(c);
        }
    }
    Err("No number found".to_string())
}

fn last_number(s: &str) -> Result<char, String> {
    for c in s.chars().rev() {
        if c.is_digit(10) {
            return Ok(c);
        }
    }
    Err("No number found".to_string())
}

fn chars_to_string(chars: Vec<char>) -> usize {
    chars.into_iter().collect::<String>().parse().unwrap()
}

fn main() -> io::Result<()> {                                                 
    let path = Path::new("../data.txt");                                      
    let file = File::open(&path)?;                                            
    let reader = io::BufReader::new(file);    

    let mut total: usize = 0;

    for line in reader.lines() {                                                                                                                         
        let line = line?;                      
        let first_num = first_number(&line).unwrap();
        let last_num = last_number(&line).unwrap();
        println!("Line: {}, First Number: {}, Last Number {}", line, first_num, last_num);                 
        let num = chars_to_string(vec![first_num, last_num]);
        println!("Number: {}", num);
        total += num;
    }                                                                                       
    println!("Total: {}", total);
    Ok(())                                                                    
}       
