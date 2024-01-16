
extern crate rand;

use std::{io, cmp::Ordering};
use rand::Rng;


fn main() {
    println!("Guess the Number!");
    // uses rand to generate a random secret_num
    let secret_num :u32 = rand::thread_rng().gen_range(1..101);

    loop {
        print!("Enter Your Guess: ");

        let mut guess = String::new();

        io::stdin().read_line(&mut guess)
            .expect("Failed while reading line.");

        let guess : u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };
        
        println!("Your guess is :{}",guess);
        
        match guess.cmp(&secret_num){
            Ordering::Less => println!("Your guess {} is low.",guess),
            Ordering::Greater => println!("Your GUESS {} is too high.",guess),
            Ordering::Equal => {
                println!("Yay...Your gueSS {} is correct,You Win!",guess);
                break;
            }
        }
    }  

}
