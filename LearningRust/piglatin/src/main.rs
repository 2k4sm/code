use std::io;

fn main() {
    let mut user_input = String::new();

    io::stdin()
        .read_line(&mut user_input)
        .expect("taking inputs failed.");

    let vowels = vec!["a", "e", "i", "o", "u"];

    let first_chr = &user_input.get(0..1).unwrap();

    if !vowels.contains(&first_chr) {
        let del_chr = user_input.remove(0);
        let new_line = user_input.pop().unwrap();
        user_input.push('-');
        user_input.push(del_chr);
        user_input.push_str("ay");
        user_input.push(new_line);
    } else {
        user_input.push_str("-hay");
        let new_line = user_input.pop().unwrap();
        user_input.push(new_line);
    }

    println!("{}", user_input);
}
