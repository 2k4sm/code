use std::io;

fn main() {
    let farhan = &mut String::new();
    println!("Enter The Temperature in farhanheit: ");

    io::stdin()
        .read_line(farhan)
        .expect("Not a floating number");

    println!();

    let farhan: f32 = farhan.trim().parse().unwrap();
    far_to_cel(&farhan);
}

fn far_to_cel(farhan: &f32) {
    println!(
        "{} farhanheit in celcius is {}",
        *farhan,
        (*farhan - 32.0) / 1.8
    );
}
