package main

import "fmt"

func main() {
	var array = []int{1, 2, 3, 4}
	array = append(array, 5)
	array = append(array, 6)

	for _, i := range array {
		fmt.Println("Address: ", &i)
		fmt.Println("Value: ", i)
	}
	fmt.Println("Array VAL:", array)
}
