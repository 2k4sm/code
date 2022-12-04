package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	sort(arr)
}

func sort(arr []string) {
	fmt.Println("Unsorted Array:", arr)
	for i := range arr {
		for k := range arr {
			if arr[i] < arr[k] {
				arr[i], arr[k] = arr[k], arr[i]
			}
		}
	}
	fmt.Println("Sorted Array:", arr)
}
