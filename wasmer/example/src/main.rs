// use std::env;
// use std::fs;
// use std::io::prelude::*;

// extern "C" {
//     // Use `js_namespace` here to bind `console.log(..)` instead of just
//     fn host_function() -> i32;
// }

// Entry point to our WASI applications
fn main() {
    // let args: Vec<String> = env::args().collect();
    // let program = args[0].clone();

    // if args.len() < 3 {
    //     eprintln!("usage: {} <target-file> <text>", program);
    //     return;
    // }

    // let mut file = fs::File::create(&args[1]).unwrap();

    // // Write the text to the file we created
    // write!(file, "Hello").unwrap();

    // unsafe {
    //     let num = host_function();
    //     println!("{}", num)
    // }

    println!("Finish run wasm.")
}
