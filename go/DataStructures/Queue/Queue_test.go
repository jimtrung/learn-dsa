package Queue

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	assert.Equal(t, 1, queue.Peek(), "Front of queue should be 1")

	queue.Enqueue(2)
	assert.Equal(t, 1, queue.Peek(), "Front of queue should remain 1 after enqueuing 2")
}

func TestDequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)

	err := queue.Dequeue()
	assert.Nil(t, err, "Dequeue should not return an error")
	assert.Equal(t, 2, queue.Peek(), "Front of queue should be 2 after dequeuing 1")

	err = queue.Dequeue()
	assert.Nil(t, err, "Dequeue should not return an error")
	assert.True(t, queue.IsEmpty(), "Queue should be empty after dequeuing all elements")

	err = queue.Dequeue()
	assert.NotNil(t, err, "Dequeue should return an error when the queue is empty")
}

func TestPeek(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(5)
	assert.Equal(t, 5, queue.Peek(), "Peek should return the front element")
}

func TestIsEmpty(t *testing.T) {
	queue := NewQueue()
	assert.True(t, queue.IsEmpty(), "New queue should be empty")

	queue.Enqueue(1)
	assert.False(t, queue.IsEmpty(), "Queue should not be empty after enqueuing an element")

	queue.Dequeue()
	assert.True(t, queue.IsEmpty(), "Queue should be empty after dequeuing all elements")
}

