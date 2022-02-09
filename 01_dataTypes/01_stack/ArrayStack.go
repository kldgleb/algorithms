package stack

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	slice []string
}

//func main() {
//	s := &stack{[]string{}}
//	s.push("1")
//	s.push("2")
//	s.push("3")
//
//	err := s.pop()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println("Size: ", s.size())
//	fmt.Println(s.slice)
//}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{[]string{}}
}

func (s *ArrayStack) Push(val string) {
	s.slice = append(s.slice, val)
}

func (s *ArrayStack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("empty stack")
	}
	buffer := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return buffer, nil
}

func (s *ArrayStack) Size() int {
	return len(s.slice)
}

func (s *ArrayStack) IsEmpty() bool {
	return s.slice == nil
}

func (s *ArrayStack) Top() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("empty stack")
	}
	return s.slice[len(s.slice)-1], nil
}

func (s *ArrayStack) Iterate() (result string) {
	for i, v := range s.slice {
		result += fmt.Sprintf("Stack index: %d value: %s \n", i, v)
	}
	return result
}
