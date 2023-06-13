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
func (l *LinkedList) printdata() {
	toprint := l.head
	for l.length != 0 {
		fmt.Println(toprint.data)
		toprint = toprint.next
		l.length--
	}
}

func (l *LinkedList) insertList(list []int) {
	for _, i := range list {
		l.insertEnd(&Node{data: i})
	}
}

func (l *LinkedList) insertAt(index int, data int) {
	ptr := l.head
	if index > l.length || index < 0 {
		panic("Index out of range.")
	} else {
		for index != 0 {
			ptr = ptr.next
			index--
		}
		changeval := ptr.next
		ptr.next = &Node{data: data}
		ptr.next.next = changeval
		l.length++
	}
}

func main() {
	llist := LinkedList{}
	llist2 := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 50}
	node3 := &Node{data: 31}
	node4 := &Node{data: 26}
	node5 := &Node{data: 27}
	node6 := &Node{data: 33}
	llist.insertEnd(node1)
	llist.insertEnd(node2)
	llist.insertEnd(node3)
	llist.insertEnd(node4)
	llist.insertEnd(node5)
	llist.insertEnd(node6)
	llist.printdata()
	list := []int{1, 2, 3, 4, 5, 6}
	llist2.insertList(list)
	llist2.insertAt(3, 9)
	llist2.insertAt(4, 10)
	llist2.insertAt(5, 12)
	llist2.insertAt(6, 12)
	llist2.insertAt(3, 9)
	llist2.insertAt(3, 9)
	llist2.insertAt(3, 9)
	llist2.insertAt(10, 15)

	llist2.printdata()

}
