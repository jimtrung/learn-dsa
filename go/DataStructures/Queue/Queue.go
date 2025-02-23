package Queue

import "errors"

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

type Queue struct {
    head *Node
    tail *Node
}

func NewQueue() *Queue {
    return &Queue{
        head: nil,
        tail: nil,
    }
}

func (q *Queue) Enqueue(value int) {
    newNode := NewNode(value)

    if q.tail == nil {
        q.head = newNode
        q.tail = newNode
    } else {
        q.tail.next = newNode
        q.tail = newNode
    }
}

func (q *Queue) Dequeue() error {
    if q.IsEmpty() {
        return errors.New("Empty queue. Can not perform Dequeue")
    }
    if q.head.next == nil {
        q.head = nil
        q.tail = nil
    } else {
        q.head = q.head.next
    }

    return nil
}

func (q *Queue) Peek() int {
    return q.head.data
}

func (q *Queue) IsEmpty() bool {
    return q.head == nil
}
