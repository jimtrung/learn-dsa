package Stack

import (
	"errors"
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

type Stack struct {
	head *Node
}

func NewStack() *Stack {
	return &Stack{
		head: nil,
	}
}

func (s *Stack) Push(value int) {
	newNode := NewNode(value)

	if s.IsEmpty() {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return errors.New("Empty stack. Can not perform Pop")
	}
	s.head = s.head.next

	return nil
}

func (s *Stack) Peek() int {
    if s.IsEmpty() {
        return -1
    }

    return s.head.data
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
