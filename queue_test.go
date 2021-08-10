package main

import (
	"errors"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPush(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)

	assert.Equal(t, queue.Peek(), 1)

	queue.Push(2)

	assert.Equal(t, queue.Peek(), 1)
}

func TestPush2(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	_, _ = queue.Pop()
	assert.Equal(t, queue.Peek(), 2)
}

func TestPopEmptyError(t *testing.T) {
	queue := NewQueue()
	_, err := queue.Pop()
	assert.Equal(t, err, errors.New(QueueIsEmptyError))
}

func TestPop(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)

	value, err := queue.Pop()
	assert.Equal(t, err, nil)
	assert.Equal(t, value, 1)
}
