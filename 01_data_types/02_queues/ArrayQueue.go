package main

import (
	"errors"
	"fmt"
)

type queue struct {
	slice []string
}

func main() {
	q := &queue{[]string{}}
	q.push("1")
	q.push("2")
	q.push("3")
	err := q.pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q.slice)
}

func (q *queue) push(val string) {
	q.slice = append(q.slice, val)
}

func (q *queue) pop() error {
	if q.isEmpty() {
		return errors.New("queue is empty")
	}
	q.slice = q.slice[len(q.slice)-(len(q.slice)-1):]
	return nil
}

func (q *queue) isEmpty() bool {
	return q.slice == nil
}

func (q *queue) size() int {
	return len(q.slice)
}
