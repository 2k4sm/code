package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) insertBegin(n *Node) {
	nextval := l.head
	l.head = n
	l.head.next = nextval
	l.length++
}

func (l *LinkedList) insertEnd(n *Node) {
	if l.length == 0 {
		l.head = n
		l.length++
	} else {
		ptr := l.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = n
		l.length++
	}

}

func (l *LinkedList) insertList(list []int) {
	for _, i := range list {
		l.insertEnd(&Node{data: i})
		l.length++
	}
}

func (l *LinkedList) insertAt(index int, data int) {
	ptr := l.head
	if index > l.length || index < 0 {
		panic("Index out of range.")
	} else {
		for index-1 != 0 {
			ptr = ptr.next
			index--
		}
		changeval := ptr.next
		ptr.next = &Node{data: data}
		ptr.next.next = changeval
		l.length++
	}
}

func (l *LinkedList) printData() {
	ptr := l.head
	for ptr != nil {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}

func (l *LinkedList) insertAfter(valueafter int, data int) {
	ptr := l.head

	for valueafter != ptr.data {
		ptr = ptr.next
	}
	changeval := ptr.next
	ptr.next = &Node{data: data}
	ptr.next.next = changeval
	l.length++

}

func (l *LinkedList) RemoveAt(index int) {
	currPtr := l.head
	prevPtr := currPtr
	nextptr := l.head.next

	for index != 0 {
		prevPtr = currPtr
		currPtr = currPtr.next
		nextptr = nextptr.next
		index--
	}
	prevPtr.next = nextptr
	currPtr = nil

}

func main() {
	llist := LinkedList{}

	list := []int{1, 2, 3, 4, 5, 6}
	llist.insertList(list)
	llist.RemoveAt(3)
	llist.RemoveAt(2)
	llist.printData()

}
