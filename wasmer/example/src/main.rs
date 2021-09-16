use std::env;
use std::fs;
use std::io::prelude::*;

// Entry point to our WASI applications
fn main() {
    let args: Vec<String> = env::args().collect();
    let program = args[0].clone();

    if args.len() < 3 {
        eprintln!("usage: {} <target-file> <text>", program);
        return;
    }

    let mut file = fs::File::create(&args[1]).unwrap();

    // Write the text to the file we created
    write!(file, "Hello").unwrap();
}
