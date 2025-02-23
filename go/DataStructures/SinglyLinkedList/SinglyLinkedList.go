package SinglyLinkedList

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
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
	}
}

func (list *LinkedList) PushFront(value int) {
    newNode := NewNode(value)

    if list.head == nil {
        list.head = newNode
        list.tail = newNode
    } else {
        newNode.next = list.head
        list.head = newNode
    }
}

func (list *LinkedList) PushBack(value int) {
    newNode := NewNode(value)

    if list.tail == nil {
        list.head = newNode
        list.tail = newNode
    } else {
        list.tail.next = newNode
        list.tail = newNode
    }
}

func (list *LinkedList) PopFront() error {
    if list.IsEmpty() {
        return errors.New("Empty list. Can not perform PopFront")
    }

    if list.head.next == nil {
        list.head = nil
        list.tail = nil
    } else {
        list.head = list.head.next
    }

    return nil
}

func (list *LinkedList) PopBack() error {
    if list.IsEmpty() {
        return errors.New("Empty list. Can not perform PopBack")
    }

    if list.head.next == nil {
        list.head = nil
        list.tail = nil
    } else {
        curr := list.head

        for curr.next.next != nil {
            curr = curr.next
        }
        curr.next = nil
        list.tail = curr;
    }

    return nil
}

func (list *LinkedList) IsEmpty() bool {
    return list.head == nil
}

func (list *LinkedList) Display() {
    curr := list.head

    for curr != nil {
        fmt.Printf("%d->", curr.data)
        curr = curr.next
    }
    fmt.Println("NULL")
}
