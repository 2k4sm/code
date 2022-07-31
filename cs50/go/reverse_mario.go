package main

import (
	"fmt"
)

func main() {
	//define the required variables
	var (
		length int
		i      int = 1
		dots   int = length
	)

	//taking the length of rev_pyramid as a user input
	fmt.Printf("Enter the length of the pyramid(1-8):")
	fmt.Scanf("%d", &length)
	//check that the length stays between (1-8).
	for length < 1 || length > 8 {
		fmt.Printf("Re-enter the length from the given range(1-8):")
		fmt.Scanf("%d", &length)
	}
	//define a logic to print THE pattern.
	for i <= length {
		j := 1
		//how many dots to make:
		dots = length - i
		//how many hash to make:
		hash := length - dots
		for j <= i {
			//To make the dots
			dots -= 1
			for dots >= 0 {
				fmt.Printf(" ")
				dots--
			}

			//To make the hashes.
			hash -= 1
			for hash >= 0 {
				fmt.Printf("#")
				hash--
			}

			j++
		}
		i++
		fmt.Println()

	}
}
