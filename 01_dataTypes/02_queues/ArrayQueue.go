package queue

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	Slice []int
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{[]int{}}
}

func (q *ArrayQueue) Enqueue(val int) {
	q.Slice = append(q.Slice, val)
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	buffer := q.Slice[0]
	q.Slice = q.Slice[len(q.Slice)-(len(q.Slice)-1):]
	return buffer, nil
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.Slice == nil || q.Size() == 0
}

func (q *ArrayQueue) Size() int {
	return len(q.Slice)
}

func (q *ArrayQueue) Iterate() (result string) {
	for i, v := range q.Slice {
		result += fmt.Sprintf("Queue index: %d value: %d \n", i, v)
	}
	return result
}
