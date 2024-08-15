package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insertFront(data int) {
	var node *Node = &Node{
		data: data,
		next: nil,
	}

	if l.head == nil {
		l.head = node
		return
	}
	node.next = l.head
	l.head = node
}

func (l *LinkedList) insertEnd(data int) {
	var node *Node = &Node{
		data: data,
		next: nil,
	}

	if l.head == nil {
		l.head = node
		return
	}

	var temp *Node = l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

func (l *LinkedList) display() {
	var temp *Node = l.head
	fmt.Println("LinkedList:")
	for temp != nil {
		fmt.Printf("%d-> ", temp.data)
		temp = temp.next
	}
}

func main() {
	ll := &LinkedList{}
	ll.insertFront(10)
	ll.insertFront(20)
	ll.insertFront(30)
	ll.insertEnd(40)
	ll.display()
}
