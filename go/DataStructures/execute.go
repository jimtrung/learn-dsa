package main

import (
	SinglyLinkedList "github.com/jimtrung/dsa-practice/go/data-structures/SinglyLinkedList"
)

func main() {
	list := SinglyLinkedList.NewLinkedList()

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)

    list.Display()
}
