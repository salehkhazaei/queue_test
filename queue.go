package main

import "errors"

const QueueIsEmptyError = "queue is empty"
type Queue struct {
	list LinkedList
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Pop() (interface{}, error) {
	v, err := q.list.RemoveFirst()

	if err != nil && err.Error() == LinkedListEmptyError {
		return nil, errors.New(QueueIsEmptyError)
	}

	return v, nil
}

func (q *Queue) Push(value interface{}) {
	q.list.PutLast(value)
}

func (q *Queue) Peek() interface{} {
	return q.list.GetFirst()
}
