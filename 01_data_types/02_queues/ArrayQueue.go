package queue

import (
	"errors"
	"fmt"
)

type arrayQueue struct {
	slice []string
}

func NewArrayQueue() *arrayQueue {
	return &arrayQueue{[]string{}}
}

func (q *arrayQueue) Enqueue(val string) {
	q.slice = append(q.slice, val)
}

func (q *arrayQueue) Dequeue() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("queue is empty")
	}
	buffer := q.slice[0]
	q.slice = q.slice[len(q.slice)-(len(q.slice)-1):]
	return buffer, nil
}

func (q *arrayQueue) IsEmpty() bool {
	return q.slice == nil
}

func (q *arrayQueue) Size() int {
	return len(q.slice)
}

func (q *arrayQueue) Iterate() {
	for i, v := range q.slice {
		fmt.Printf("Queue index: %d value: %s \n", i, v)
	}
}
