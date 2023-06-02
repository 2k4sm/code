package temptable

import (
	"fmt"
	conv "temptables/conv"
)

func DrawTable() {

	tablen1 := 140
	tablen2 := 140

	for tablen1 > 0 {
		if tablen1 == 140/2 {
			fmt.Printf("|C|F|")
		} else if tablen1 == 1 {
			fmt.Printf("|")
		} else {
			fmt.Printf("=")
		}
		tablen1--
	}
	fmt.Println()
	x := -40
	for x < 100 {
		v := conv.Celsius(x)
		fmt.Printf("|%d|%.4v|\n", x, v.CelConv())
		x++

	}
	for tablen2 > 0 {
		if tablen2 == 140/2 {
			fmt.Printf("|F|C|")
		} else if tablen2 == 1 {
			fmt.Printf("|")
		} else {
			fmt.Printf("=")
		}
		tablen2--
	}
	fmt.Println()
	x = -40
	for x < 100 {
		v := conv.Fahrenheit(x)
		fmt.Printf("|%d|%.4v|\n", x, v.FarhenConv())
		x++

	}
}
