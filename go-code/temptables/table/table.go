package temptable

import (
	"fmt"
	conv "temptables/conv"
)


func DrawTable(Myfunc conv.Temperature) {

	tablen := 140

	for tablen > 0 {
		if tablen == 140/2 {
			fmt.Printf("|C|F|")
		} else if tablen == 1 {
			fmt.Printf("|")
		} else {
			fmt.Printf("=")
		}
		tablen--
	}
	fmt.Println()
	x := -40
	for x < 100 {
		conv.K =float32(x)
		fmt.Printf("|%d|%.4v|\n", x, conv.K)
		x++

	}
}
