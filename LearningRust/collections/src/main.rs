use std::{collections::HashMap, io, u32};

fn main() {
    let mut int_list: Vec<u32> = Vec::new();
    populate(&mut int_list);

    let list_len = int_list.len();

    print_median(&int_list, &list_len);

    print_mean(&int_list, &list_len);

    let mode = print_mode(&int_list);

    println!("mode is: {} with value: {}", mode.0, mode.1);
}

fn print_mode(int_list: &Vec<u32>) -> (u32, u32) {
    let mut mode_cont: HashMap<u32, u32> = HashMap::new();

    for val in int_list {
        let mode_cnt = mode_cont.entry(*val).or_insert(0 as u32);
        *mode_cnt += 1;
    }

    let mut mode: (u32, u32) = (0, 0);
    for (key, val) in &mode_cont {
        if mode_cont.get(&key).unwrap() >= &mode.1 {
            mode = (*key, *val);
        }
    }
    mode
}

fn print_median(int_list: &Vec<u32>, list_len: &usize) {
    let default: u32 = 0;
    let mut sorted_val = int_list.clone();
    sorted_val.sort();
    println!(
        "median is: {}",
        &sorted_val.get(*list_len / 2).unwrap_or(&default)
    );
}

fn print_mean(int_list: &Vec<u32>, list_len: &usize) {
    let mut mean = 0;

    for val in int_list {
        mean += *val;
    }

    println!("mean is: {}", mean / *list_len as u32);
}
fn populate(int_list: &mut Vec<u32>) {
    loop {
        let mut input_int = String::new();

        io::stdin()
            .read_line(&mut input_int)
            .expect("failed reading from standardinput.");

        if input_int.eq("\n") {
            break;
        }

        let input_int: u32 = input_int.trim().parse().expect("invalid input.");

        int_list.push(input_int);
    }
}
