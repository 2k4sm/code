package main

import (
	"fmt"

	T "./temperature"
)

func main() {
	fmt.Println(Ktoc(233))
	fmt.Println(T.Ctof(Ktoc(233)))
	fmt.Println(T.Ktof(0))
}
func Ktoc(temp int) int {
	const kelv0 int = 273

	return temp - kelv0

}
