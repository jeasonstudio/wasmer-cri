#[no_mangle]
pub extern "C" fn sum(x: i32, y: i32) -> i32 {
    x + y
}
pub extern "C" fn _start(x: i32, y: i32) -> i32 {
    x * y
}

fn main() {
    println!("simple example");
}
