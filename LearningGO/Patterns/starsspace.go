package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	plength, _ := strconv.Atoi(os.Args[1])
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
