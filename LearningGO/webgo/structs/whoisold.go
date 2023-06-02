package main

import "fmt"

type person struct {
	name string
	age  int
}

func whold(P1, P2 person) (person, person) {

	if P1.age > P2.age {
		return P1, P2
	} else {
		return P2, P1
	}
}

func main() {

	P := person{"Papun", 19}
	B := person{"Bapun", 28}

	old, notold := whold(P, B)
	agediff := (old.age - notold.age)
	fmt.Printf("%s is older than %s by %d years.\n", old.name, notold.name, agediff)

}
