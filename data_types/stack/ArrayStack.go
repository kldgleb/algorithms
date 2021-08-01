package main

import (
	"errors"
	"fmt"
)

type stack struct {
	slice []string
}

func main() {
	s := &stack{[]string{}}
	s.push("1")
	s.push("2")
	s.push("3")

	err := s.pop()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Size: ", s.size())
	fmt.Println(s.slice)
}

func (s *stack) push(val string) {
	s.slice = append(s.slice, val)
}

func (s *stack) pop() error {
	if s.isEmpty() {
		return errors.New("empty stack")
	}
	s.slice = s.slice[:len(s.slice)-1]
	return nil
}

func (s *stack) size() int {
	return len(s.slice)
}

func (s *stack) isEmpty() bool {
	return s.slice == nil
}
