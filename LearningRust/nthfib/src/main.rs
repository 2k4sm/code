use std::io;

fn main() {
    let n = &mut String::new();

    io::stdin().read_line(n).expect("Error Reading n");

    let n: i32 = n.trim().parse().unwrap();

    println!("{}th fib number is {}", n, fib(n));
}

fn fib(n: i32) -> i32 {
    if n == 1 {
        return 1;
    }

    if n == 0 {
        return 0;
    }

    fib(n - 1) + fib(n - 2)
}
