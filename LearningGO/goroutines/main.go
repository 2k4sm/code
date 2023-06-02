package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

// channel

func sum(a []int, c chan int) {
	t := 0
	for _, v := range a {
		t += v
	}
	c <- t
}

// Buffered Channels

func buffchan() {
	c := make(chan int, 3)

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)

}

// goroutines

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

// fibonacci sequence.

func fibonacci(n int, ch chan int) {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y

	}
	close(ch)
}

//using select for fibonacci.

func sfib(ch, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return

		}

	}

}
