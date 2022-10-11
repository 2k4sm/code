package main

import "fmt"

func main() {
	var command string
	fmt.Printf("Enter the command:")
	fmt.Scanln(&command)
	//fmt.Println(command)
	if command == "goeast" {
		fmt.Println("Head further up to mountains.")

	} else if command == "goinside" {
		fmt.Println("Live in cave for rest of the life...")
	} else {
		fmt.Println("Didn't quite get that.")
	}
}
