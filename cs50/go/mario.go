package main

import "fmt"

func main() {
	//Make an integer to take the length of pyramind and keep it between(1--8):
	var length int
	var i int = 1
	fmt.Printf("Enter the length of the pyramid(1-8):")
	fmt.Scanf("%d", &length)

	//Define a logic to length of pyramid is between (1-8)
	for length < 1 || length > 8 {
		fmt.Printf("Re-enter the length of the pyramid from the specified input(1-8):")
		fmt.Scanf("%d", &length)

	}
	//logic to print the pattern
	for i <= length {
		j := 1
		for j <= i {
			fmt.Printf("#")
			j++
		}
		i++
		fmt.Println()

	}

}
