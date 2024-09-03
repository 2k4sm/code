package main

import "fmt"

type E interface {
	int | string | float32
}

func main() {
	intArr := []int{2, 3, 4, 5, 6}
	stringArr := []string{"hello", "world", "why", "are", "you", "so", "difficult"}
	floatArr := []float32{3.14, 0.89, 1.23, 5.67}

	Ranger(intArr)
	Ranger(stringArr)
	Ranger(floatArr)
}

func Ranger[S E](s []S) {
	for _, val := range s {
		fmt.Println(val)
	}
}
