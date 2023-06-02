package main

import (
	"fmt"
)

var tempTyp string
var k temperature

func main() {
	fmt.Printf("Temperature_Type:")
	fmt.Scanf("%v\n", &tempTyp)
	fmt.Printf("Temperature:")
	fmt.Scanf("%d\n", &k)
	fmt.Println("Celsius:", k.celsius())

}

type (
	temperature float32
)

func (k temperature) celsius() temperature {
	if tempTyp == "farhenheit" {
		cel1 := (k - 32) * (5 / 9)
		return cel1
	} else if tempTyp == "kelvin" {
		cel2 := k - 273.15
		return cel2
	} else {
		return k
	}
}
