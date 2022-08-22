package queue

type Queue interface {
	Enqueue(int)
	Dequeue() (int, error)
	IsEmpty() bool
	Size() int
	Iterate() string
}

func Enqueue(q Queue, val int) {
	q.Enqueue(val)
}

func Dequeue(q Queue) (int, error) {
	return q.Dequeue()
}

func IsEmpty(q Queue) bool {
	return q.IsEmpty()
}

func Size(q Queue) int {
	return q.Size()
}

func Iterate(q Queue) string {
	return q.Iterate()
}
