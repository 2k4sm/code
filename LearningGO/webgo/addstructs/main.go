package main

import "fmt"

type computer struct {
	cpu   string
	tdp   int
	os    string
	phone int
}

type robot struct {
	computer
	utility string
	phone   int
}

func main() {
	sm := robot{computer{"intel", 50, "gnu/linux", 1234}, "Home", 5467}

	fmt.Printf("%s cpu is used with %d watt tdp and %s os for %s use.%d\n", sm.cpu, sm.tdp, sm.os, sm.utility, sm.phone)

}
