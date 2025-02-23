package CircularLinkedList

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewNode(value int) *Node {
	return &Node{
		data: value,
		next: nil,
	}
}

type LinkedList struct {
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		tail: nil,
	}
}

func (list *LinkedList) PushFront(value int) {
	newNode := NewNode(value)

	if list.IsEmpty() {
		newNode.next = newNode
		list.tail = newNode
	} else {
		newNode.next = list.tail.next
		list.tail.next = newNode
	}
}

func (list *LinkedList) PushBack(value int) {
	newNode := NewNode(value)

	if list.IsEmpty() {
		newNode.next = newNode
		list.tail = newNode
	} else {
		newNode.next = list.tail.next
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *LinkedList) PopFront() error {
	if list.IsEmpty() {
		return errors.New("Empty list. Failed to perform PopFront")
	}
	if list.tail.next == list.tail {
		list.tail = nil
	} else {
		list.tail.next = list.tail.next.next
	}

	return nil
}

func (list *LinkedList) PopBack() error {
	if list.IsEmpty() {
		return errors.New("Empty list. Failed to perform PopBack")
	}
	if list.tail.next == list.tail {
		list.tail = nil
	} else {
		curr := list.tail.next

		for curr.next != list.tail {
			curr = curr.next
		}
		curr.next = curr.next.next
		list.tail = curr
	}

	return nil
}

func (list *LinkedList) IsEmpty() bool {
	return list.tail == nil
}

func (list *LinkedList) Display() {
	if !list.IsEmpty() {
		curr := list.tail.next
		lastNode := list.tail

		for curr != list.tail {
			fmt.Printf("%d->", curr.data)
			curr = curr.next
		}

		fmt.Printf("%d->", lastNode.data)
	}
	fmt.Println("NULL")
}
