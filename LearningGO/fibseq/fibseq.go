package main

import (
	"fmt"
)

func main(){
	fmt.Println(fibseq(4))
}

func fibseq(num int) int{
	if (num == 0 || num == 1){
		return num
	}
	return fibseq(num-1) + fibseq(num-2)
}
