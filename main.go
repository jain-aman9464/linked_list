package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	len  int
	head *Node
	tail *Node
}

func (l *LinkedList) insertAtBeg(value int) {
	newNode := &Node{
		value: value,
	}

	if l.len == 0 {
		l.head = newNode
		l.tail = newNode
		l.len++
		return
	}

	newNode.next = l.head
	l.head = newNode
	l.len++
}

func (l *LinkedList) insertAtEnd(value int) {
	newNode := &Node{
		value: value,
	}

	if l.len == 0 {
		l.head = newNode
		l.tail = newNode
		l.len++
		return
	}

	//curr := l.head
	//for curr.next != nil {
	//	curr = curr.next
	//}

	curr := l.tail
	curr.next = newNode
	l.tail = newNode
	l.len++
}

func (l *LinkedList) display() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func main() {
	l := &LinkedList{}
	l.insertAtEnd(1)
	l.insertAtEnd(2)
	l.insertAtEnd(3)

	l.display()
}
