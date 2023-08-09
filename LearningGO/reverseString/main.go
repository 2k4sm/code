package main

import (
	ms "example/sm2k4/reverseString/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(ms.ReverseRunes("Hello,world!"))
	fmt.Println(cmp.Diff("Hello World!", "Hello Go!"))
}
