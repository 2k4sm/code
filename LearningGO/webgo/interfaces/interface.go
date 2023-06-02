package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

type Men interface {
	Sayhi()
	Sing(lyrics string)
}

func (h Human) Sayhi() {
	fmt.Printf("HII,I am %s You can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Printf("%s\n", lyrics)

}

func (e Employee) Sayhi() {
	fmt.Printf("Hii, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}

	i := make([]Men, 3)

	i[0], i[1], i[2] = paul, sam, mike

	for _, j := range i {
		j.Sayhi()
		j.Sing("laaalalalllalaaa....")
	}

}
