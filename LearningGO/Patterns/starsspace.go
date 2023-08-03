package main

import "fmt"

func main() {

	plength := 10
	stars := 1
	for plength != 0 {
		for i := 0; i < stars; i++ {
			fmt.Print("*")
		}
		stars++
		fmt.Println()
		plength--
	}

}
