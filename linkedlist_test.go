package main

import (
	"errors"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPutLast(t *testing.T) {
	l := NewLinkedList()
	l.PutLast(1)
	assert.Equal(t, l.GetLast(), 1)

	l.PutLast(2)
	assert.Equal(t, l.GetLast(), 2)
}

func TestRemoveFirstLinkedListEmpty(t *testing.T) {
	l := NewLinkedList()
	_, err := l.RemoveFirst()
	assert.Equal(t, err, errors.New(LinkedListEmptyError))
}

func TestRemoveFirstOneItem(t *testing.T) {
	l := NewLinkedList()

	l.PutLast(1)
	_, err := l.RemoveFirst()
	assert.Equal(t, err, nil)
	assert.Equal(t, l.GetLast(), 1)
}

func TestRemoveFirstTwoItems(t *testing.T) {
	l := NewLinkedList()

	l.PutLast(1)
	l.PutLast(2)

	value, err := l.RemoveFirst()
	assert.Equal(t, err, nil)
	assert.Equal(t, value, 1)

	value, err = l.RemoveFirst()
	assert.Equal(t, err, nil)
	assert.Equal(t, value, 2)
}

func TestGetFirst(t *testing.T) {
	l := NewLinkedList()

	l.PutLast(1)
	l.PutLast(2)

	assert.Equal(t, l.GetFirst(), 1)
}
