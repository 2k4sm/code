package main

import (
	"fmt"
)

func Quarters(cents int) int {
	return cents / 25
}
func Dimes(cents int) int {
	return cents / 10
}
func Nickels(cents int) int {
	return cents / 5
}
func Penny(cents int) int {
	return cents
}

func main() {
	var (
		cent  int
		coins int = 0
	)
	fmt.Printf("Enter the number of cents:")
	fmt.Scanf("%d", &cent)
	for cent < 0 {
		fmt.Printf("Enter the cents in +ve range:")
		fmt.Scanf("%d", &cent)
	}
	for cent > 0 {
		switch cent != 0 {
		case cent >= 25:
			coins += Quarters(cent)
			cent -= Quarters(cent) * 25
		case cent < 25 && cent >= 10:
			coins += Dimes(cent)
			cent -= Dimes(cent) * 10
		case cent < 10 && cent >= 5:
			coins += Nickels(cent)
			cent -= Nickels(cent) * 5
		default:
			coins += Penny(cent)
			cent -= Penny(cent) * 1
		}

	}
	fmt.Printf("The total coins spared is %d\n", coins)

}
