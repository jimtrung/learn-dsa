package DoublyLinkedList

import (
	"errors"
	"fmt"
)

type Node struct {
    data int
    next *Node
    prev *Node
}

func NewNode(value int) *Node {
    return &Node{
        data: value,
        next: nil,
        prev: nil,
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
        list.head.prev = newNode
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
        newNode.prev = list.tail
        list.tail = newNode
    }
}

func (list *LinkedList) PopFront() error {
    if list.IsEmpty() {
        return errors.New("Empty list. Cannot perform PopFront")
    }
    if list.head.next == nil {
        list.head = nil
        list.tail = nil
    } else {
        list.head = list.head.next
        list.head.prev = nil
    }

    return nil
}

func (list *LinkedList) PopBack() error {
    if list.IsEmpty() {
        return errors.New("Empty list. Cannot perform PopBack")
    }
    if list.head.next == nil {
        list.head = nil
        list.tail = nil
    } else {
        list.tail = list.tail.prev
        list.tail.next = nil
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
