package main

import "fmt"

type Node struct {
	data string
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) PopulateString(input string) {
	currPtr := l.head
	for _, j := range input {
		for currPtr != nil {
			currPtr = currPtr.next
		}
		currPtr = &Node{data: string(j)}
	}
}

func main() {
	llist := LinkedList{}
	llist.PopulateString("apple")
	fmt.Println(llist.head)
}
