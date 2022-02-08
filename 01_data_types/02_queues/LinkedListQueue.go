package queue

import (
	"errors"
	"fmt"
)

type linkedListQueue struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	val  string
	next *node
}

func NewLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{nil, nil, 0}
}

func (q *linkedListQueue) Enqueue(val string) {
	oldTail := q.tail
	node := &node{val, nil}
	q.tail = node
	if q.IsEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.length++
}

func (q *linkedListQueue) Dequeue() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("empty queue")
	}
	toDelete := q.head.val
	q.head = q.head.next
	q.length--
	return toDelete, nil
}

func (q *linkedListQueue) IsEmpty() bool {
	return q.head == nil
}

func (q *linkedListQueue) Size() int {
	return q.length
}

func (q linkedListQueue) Iterate() {
	index := 0
	for q.head != nil {
		fmt.Printf("Queue index: %d value: %s \n", index, q.head.val)
		q.head = q.head.next
		index++
	}
}
