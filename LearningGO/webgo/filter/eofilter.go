package main

import "fmt"

type intfunc func(int) bool

func isOdd(integer int) bool {
	return integer%2 != 0

}
func isEven(integer int) bool {
	return integer%2 == 0
}
func filter(slice []int, oeFilter intfunc) []int {
	var result []int
	for _, i := range slice {
		if oeFilter(i) {
			result = append(result, i)
		}
	}
	return result
}
func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 11, 56, 31, 12}

	sliceOdd := filter(slice, isOdd)
	sliceEven := filter(slice, isEven)
	fmt.Printf("Even: %d\n Odd: %d\n", sliceEven, sliceOdd)
}
