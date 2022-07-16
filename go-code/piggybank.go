package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nickels := 0.05
	dimes := 0.10
	quarters := 0.25
	bank := 0.0
	fmt.Printf("Your initial bank balance:%.4v \n", bank)
	for bank < 20.0 {
		random := rand.Intn(3) + 1
		switch random {
		case 1.0:
			bank += nickels
			//fmt.Printf("i\n")
		case 2.0:
			bank += dimes
			//fmt.Printf("j\n")
		case 3.0:
			bank += quarters
			//fmt.Printf("K\n")
		default:
			fmt.Println("Error")
		}
		fmt.Printf("Your current balance:%.4v \n", bank)

	}
	fmt.Printf("Your bank balance:%.4v \n", bank)
}
