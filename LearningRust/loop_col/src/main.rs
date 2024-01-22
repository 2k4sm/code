fn main() {
    let collection = vec![10, 11, 12, 14, 16, 21, 3, 64, 7, 8, 9, 10];
    loop_col(collection);
}

fn loop_col(col: Vec<u32>) {
    let mut _index: u32 = 0;

    for val in col.iter() {
        println!("Current Value in collection : {}", val);
        _index += 1;
    }
    println!("Total elements in collection : {}", col.len());
}
