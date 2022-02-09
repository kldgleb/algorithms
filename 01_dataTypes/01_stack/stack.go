package stack

type Stack interface {
	Push(val string)
	Pop() (string, error)
	Size() int
	Top() (string, error)
	IsEmpty() bool
	Iterate() string
}

func Push(stack Stack, val string) {
	stack.Push(val)
}

func Pop(stack Stack) (string, error) {
	return stack.Pop()
}

func Size(stack Stack) int {
	return stack.Size()
}

func Top(stack Stack) (string, error) {
	return stack.Top()
}

func IsEmpty(stack Stack) bool {
	return stack.IsEmpty()
}

func Iterate(stack Stack) string {
	return stack.Iterate()
}
