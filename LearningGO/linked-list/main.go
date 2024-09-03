package main

import "fmt"

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head *element[T]
	tail *element[T]
}

func (list *List[T]) PushFront(val T) {
	if list.head == nil {
		list.head = &element[T]{val: val}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: val}
		list.tail = list.tail.next
	}
}

func (list *List[T]) collect() []T {
	var result []T
	for e := list.head; e != nil; e = e.next {
		result = append(result, e.val)
	}

	return result
}

func main() {
	strList := List[string]{}
	intList := List[int]{}
	floatList := List[float64]{}

	strList.PushFront("Hello")
	strList.PushFront("World")
	strList.PushFront("!")

	intList.PushFront(1)
	intList.PushFront(2)
	intList.PushFront(3)

	floatList.PushFront(1.1)
	floatList.PushFront(2.2)
	floatList.PushFront(3.3)

	fmt.Println(strList.collect())
	fmt.Println(intList.collect())
	fmt.Println(floatList.collect())

}
