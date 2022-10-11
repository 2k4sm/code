package main

import (
	"fmt"
)

func main() {
	fmt.Println(Ktoc(273))
	fmt.Println(Ctof(Ktoc(273)))
	fmt.Println(Ktof(0))
}
func Ktoc(temp int) int {
	const kelv0 int = 273

	return temp - kelv0

}
func Ctof(temp int) int {
	f := (temp * 9.0 / 5.0) + 32.0
	return f
}
func Ktof(temp int) int {
	c := temp - 273

	f := (c * 9.0 / 5.0) + 32.0

	return f
}
