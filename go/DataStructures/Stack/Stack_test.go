package Stack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	assert.Equal(t, 1, stack.Peek(), "Top of stack should be 1")

	stack.Push(2)
	assert.Equal(t, 2, stack.Peek(), "Top of stack should be 2")
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)

	err := stack.Pop()
	assert.Nil(t, err, "Pop should not return an error")
	assert.Equal(t, 1, stack.Peek(), "Top of stack should be 1 after popping")

	err = stack.Pop()
	assert.Nil(t, err, "Pop should not return an error")
	assert.True(t, stack.IsEmpty(), "Stack should be empty after popping all elements")

	err = stack.Pop()
	assert.NotNil(t, err, "Pop should return an error when the stack is empty")
}

func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	assert.Equal(t, 5, stack.Peek(), "Peek should return the top element")
}

func TestIsEmpty(t *testing.T) {
	stack := NewStack()
	assert.True(t, stack.IsEmpty(), "New stack should be empty")

	stack.Push(1)
	assert.False(t, stack.IsEmpty(), "Stack should not be empty after pushing an element")

	stack.Pop()
	assert.True(t, stack.IsEmpty(), "Stack should be empty after popping all elements")
}
