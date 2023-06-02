package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The searched item is at index: %d", search())
}

func search() int {
	var (
		list    []int32
		item    int32 = 0
		num     int32
		i       int32 = 0
		element int32 = 0
	)

	fmt.Printf("how many numbers in the list:")
	fmt.Scanf("%d", &num)

	for i < num {
		fmt.Printf("Enter your element:")
		fmt.Scanf("%d", &element)
		list = append(list, element)
		i++
	}
	fmt.Printf("Enter the item to search for:")
	fmt.Scanf("%d", &item)

	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}
