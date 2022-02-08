package stack

import (
	"errors"
	"fmt"
)

type LinkedListStack struct {
	head *node
	size int
}

type node struct {
	val  string
	prev *node
}

//func main() {
//	linkedList := &linkedList{nil, 0}
//
//	fmt.Println("Add A ...")
//	linkedList.push("A")
//	linkedList.itterate()
//
//	fmt.Println("Add B ...")
//	linkedList.push("B")
//	linkedList.itterate()
//
//	fmt.Println("Add C ...")
//	linkedList.push("C")
//	linkedList.itterate()
//
//	fmt.Println("Delete last element ...")
//	val, err := linkedList.pop()
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("Deleted val %v \n", val)
//	linkedList.itterate()
//
//	size := linkedList.getSize()
//	fmt.Printf("Size of list = %v \n", size)
//
//	top := linkedList.top()
//	fmt.Printf("Top value %v", top)
//}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil, 0}
}

func (l *LinkedListStack) Push(val string) {
	node := &node{val, l.head}
	l.head = node
	l.size++
}

func (l *LinkedListStack) Pop() (string, error) {
	if l.head == nil {
		return "", errors.New("stack is empty")
	}
	buffer := l.head.val
	l.head = l.head.prev
	l.size--
	return buffer, nil
}

func (l LinkedListStack) Size() int {
	return l.size
}

func (l LinkedListStack) IsEmpty() bool {
	return l.head == nil
}

func (l LinkedListStack) Top() (string, error) {
	if l.IsEmpty() {
		return "", errors.New("empty stack")
	}
	return l.head.val, nil
}

func (l LinkedListStack) Iterate() {
	index := 0
	for l.head != nil {
		fmt.Printf("Stack index: %d value: %s \n", index, l.head.val)
		index++
		l.head = l.head.prev
	}
}
