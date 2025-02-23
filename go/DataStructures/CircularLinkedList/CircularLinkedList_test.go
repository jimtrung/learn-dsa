package CircularLinkedList

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushFront(t *testing.T) {
	list := NewLinkedList()

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)

	cases := []struct {
		name     string
		actual   int
		expected int
	}{
		{"First node should be 3", list.tail.next.data, 3},
		{"Second node should be 2", list.tail.next.next.data, 2},
		{"Third node should be 1", list.tail.data, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.actual)
		})
	}
}

func TestPushBack(t *testing.T) {
	list := NewLinkedList()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	cases := []struct {
		name     string
		actual   int
		expected int
	}{
		{"First node should be 1", list.tail.next.data, 1},
		{"Second node should be 2", list.tail.next.next.data, 2},
		{"Third node should be 3", list.tail.data, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.actual)
		})
	}
}

func TestPopFront(t *testing.T) {
	list := NewLinkedList()
	list.PushBack(1)
	list.PushBack(2)

	err := list.PopFront()
	assert.Nil(t, err, "PopFront should not return an error")

	err = list.PopFront()
	assert.Nil(t, err, "PopFront should not return an error")

	err = list.PopFront()
	assert.NotNil(t, err, "PopFront should return an error when the list is empty")
}

func TestPopBack(t *testing.T) {
	list := NewLinkedList()
	list.PushBack(1)
	list.PushBack(2)

	err := list.PopBack()
	assert.Nil(t, err, "PopBack should not return an error")

	err = list.PopBack()
	assert.Nil(t, err, "PopBack should not return an error")

	err = list.PopBack()
	assert.NotNil(t, err, "PopBack should return an error when the list is empty")
}

func TestIsEmpty(t *testing.T) {
	list := NewLinkedList()
	assert.True(t, list.IsEmpty(), "New list should be empty")

	list.PushBack(1)
	assert.False(t, list.IsEmpty(), "List should not be empty after adding elements")

	list.PopFront()
	assert.True(t, list.IsEmpty(), "List should be empty after removing all elements")
}

func captureOutput(f func()) string {
    old := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    f()

    w.Close()
    var buf bytes.Buffer
    buf.ReadFrom(r)
    os.Stdout = old

    return buf.String()
}

func TestDisplay(t *testing.T) {
	list := NewLinkedList()

	output := captureOutput(func() {
		list.Display()
	})
	assert.Equal(t, "NULL\n", output, "Empty list should print 'NULL'")

	list.PushBack(1)
	list.PushBack(2)

	output = captureOutput(func() {
		list.Display()
	})
	assert.Equal(t, "1->2->NULL\n", output, "List should print '1->2->NULL'")
}

