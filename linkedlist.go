package main

import "errors"

const LinkedListEmptyError = "linkedList is empty"
type Node struct {
	value interface{}

	next *Node
}

type LinkedList struct {
	firstNode *Node
	lastNode  *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{
		firstNode: nil,
		lastNode:  nil,
	}
}

func (l *LinkedList) GetLast() interface{} {
	return l.lastNode.value
}

func (l *LinkedList) GetFirst() interface{} {
	return l.firstNode.value
}

func (l *LinkedList) PutLast(value interface{}) {
	node := Node{value: value}
	if l.firstNode == nil && l.lastNode == nil {
		l.firstNode = &node
		l.lastNode = &node
	} else {
		l.lastNode.next = &node
		l.lastNode = &node
	}
}

func (l *LinkedList) RemoveFirst() (interface{}, error) {
	if l.firstNode == nil {
		return nil, errors.New(LinkedListEmptyError)
	}

	node := l.firstNode
	l.firstNode = l.firstNode.next

	return node.value, nil
}