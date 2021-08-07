package main

import "fmt"

type linkedList struct {
	head   *node
	tail   *node
	lenght int
}
type node struct {
	val  string
	next *node
}

func main() {
	queue := &linkedList{nil, nil, 0}
	queue.isEmpty()
	queue.enqueue("1")
	queue.enqueue("2")
	queue.enqueue("3")
	queue.dequeue()
	queue.itterate()
	fmt.Println("size: ", queue.size())
}

func (q *linkedList) isEmpty() bool {
	return q.head == nil
}

func (q *linkedList) enqueue(val string) {
	oldTail := q.tail
	node := &node{val, nil}
	q.tail = node
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.lenght++
}

func (q *linkedList) dequeue() string {
	toDelete := q.head.val
	q.head = q.head.next
	if q.isEmpty() {
		q.tail = nil
	}
	q.lenght--
	return toDelete
}

func (q linkedList) itterate() {
	for q.head != nil {
		fmt.Println(q.head.val)
		q.head = q.head.next
	}
}

func (q *linkedList) size() int {
	return q.lenght
}
