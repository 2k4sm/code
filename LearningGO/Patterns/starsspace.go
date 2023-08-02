package main

import "fmt"

func main() {

	plength := 10
	stars := 1
	for plength != 0 {
		spaces := plength - stars
		for i := 1; i <= spaces; i++ {
			fmt.Print(" ")
		}
		for k := 1; k <= stars; k++ {
			fmt.Print("*")
		}
		fmt.Println()
		plength--
		stars++
	}

}
