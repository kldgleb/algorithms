package main

import (
	"errors"
	"fmt"
)

type linkedList struct {
	head *node
	size int
}

type node struct {
	val  string
	prev *node
}

func main() {
	linkedList := &linkedList{nil, 0}

	fmt.Println("Add A ...")
	linkedList.push("A")
	linkedList.itterate()

	fmt.Println("Add B ...")
	linkedList.push("B")
	linkedList.itterate()

	fmt.Println("Add C ...")
	linkedList.push("C")
	linkedList.itterate()

	fmt.Println("Delete last element ...")
	val, err := linkedList.pop()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Deleted val %v \n", val)
	linkedList.itterate()

	size := linkedList.getSize()
	fmt.Printf("Size of list = %v \n", size)
}

func (l *linkedList) push(val string) {
	node := &node{val, l.head}
	l.head = node
	l.size++
}

func (l *linkedList) pop() (string, error) {
	if l.head == nil {
		return "", errors.New("list is empty")
	}
	buffer := l.head.val
	l.head = l.head.prev
	l.size--
	return buffer, nil
}

func (l linkedList) getSize() int {
	return l.size
}

func (l linkedList) itterate() {
	for l.head != nil {
		fmt.Println(l.head.val)
		l.head = l.head.prev
	}
}
