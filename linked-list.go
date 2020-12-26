package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head

	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printList() {
	head := l.head
	for l.length > 0 {
		fmt.Printf("%d ", head.data)
		head = head.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	previousNode := l.head

	if l.length == 0 {
		return
	}

	if previousNode.data == value {
		l.head = previousNode.next
		l.length--
		return
	}

	// My answer
	/*
		for previousNode.next != nil && previousNode.next.data != value {
			previousNode = previousNode.next
		}
		if previousNode.next != nil {
			previousNode.next = previousNode.next.next
			l.length--
			return nil
		} else {
			return fmt.Errorf("No node with value %d", value)
		}
	*/

	// Hers
	for previousNode.next.data != value {
		/* after checking if the value of the next node is
		not the value to delete, check that the next of the
		currently next value is valid else assume you have
		gotten to the last node */

		if previousNode.next.next == nil {
			return
		}
		previousNode = previousNode.next
	}
	previousNode.next = previousNode.next.next
	l.length--
}

func main() {
	myList := &linkedList{}

	node1 := &node{data: 48}
	node2 := &node{data: 32}
	node3 := &node{data: 25}
	node4 := &node{data: 15}
	node5 := &node{data: 45}
	node6 := &node{data: 6}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.printList()
	myList.deleteWithValue(100)
	myList.printList()
	myList.deleteWithValue(6)
	myList.printList()

	myEmptyList := linkedList{}

	myEmptyList.deleteWithValue(9)
}
