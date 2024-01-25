fn main() {
    let strval = String::from("Hello World");

    let slice = first_word(&strval);

    println!("{}", strval);
    println!("{}", slice);
}

fn first_word(strval: &String) -> &str {
    let bytes = strval.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &strval[0..i];
        }
    }

    &strval[..]
}
