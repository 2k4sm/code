package main

import (
	"fmt"
)

func main() {
	var (
		height int
		i      int = 1
	)
	//Taking user input to get the height of pyramid
	fmt.Printf("Enter the height of the pyramid from the RANGE (1-8):")
	fmt.Scanf("%d", &height)

	//re-take the input if not in the specied range of (1-8)
	for height < 1 || height > 8 {
		fmt.Printf("Re-enter the height of the pyramid from the specified RANGE (1-8):")
		fmt.Scanf("%d", &height)
	}
	//fmt.Println(height)

	//logic to print the pyramid.
	for i <= height {

		//To print the forword dots/blanks.
		dots := height - i
		for dots > 0 {
			fmt.Printf(" ")
			dots--
		}

		//To print the left pattern.
		dash := i
		for dash > 0 {
			fmt.Printf("#")
			dash--
		}

		// To print the middle dots/blanks.
		fmt.Printf("  ")

		// To print the right pattern
		j := 1
		for j <= i {
			fmt.Printf("#")
			j++
		}
		i++
		fmt.Println()
	}
}
