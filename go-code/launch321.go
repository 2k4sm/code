package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 10
	for count > 0 {
		fmt.Println(count)
		seconds := 2
		time.Sleep(time.Duration(seconds) * time.Second)
		if rand.Intn(101) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Launch successful...")

	} else {
		fmt.Println("Launch failed...")
	}

}
