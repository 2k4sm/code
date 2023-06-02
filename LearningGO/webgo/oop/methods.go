package main

import "fmt"

const (
	white = iota
	black
	blue
	red
	yellow
)

type box struct {
	width, height, length float64
	color                 Color
}

type Color byte
type boxslice []box

//method

func (Box box) volume() float64 {
	return Box.height * Box.length * Box.width
}

//method with pointer as reciever

func (Box *box) setcolor(c Color) {
	Box.color = c
}

//method

func (boxlist boxslice) Paintblack() {
	for i, _ := range boxlist {
		boxlist[i].setcolor(black)

	}
}

//method

func (c Color) str() string {
	strs := []string{"white", "black", "blue", "red", "yellow"}

	return strs[c]
}

func main() {
	boxes := boxslice{
		box{4, 4, 4, red},
		box{10, 10, 1, yellow},
		box{1, 1, 20, black},
		box{10, 10, 1, blue},
		box{10, 30, 1, white},
		box{20, 20, 20, yellow},
	}
	fmt.Printf("No of boxes:%d\n", len(boxes))
	fmt.Printf("Vol Box1:%f\nVol Box%d:%f\n", boxes[0].volume(), len(boxes), boxes[5].volume())
	fmt.Printf("Last Box color:%v\n", boxes[5].color.str())

	boxes.Paintblack()
	fmt.Printf("New Last Box color:%v\n", boxes[5].color.str())

}
