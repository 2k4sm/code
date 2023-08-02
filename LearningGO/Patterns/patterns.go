package main

import "fmt"

func main() {
	
	plength := 10

	for plength != 0{
		i :=1
		for i = 1; i< plength;i++{
			fmt.Print(" ")
		}
		j := plength-i
		for j < plength{
			fmt.Print("*")
			j++
		}
		fmt.Println()
		plength--
	}

}
