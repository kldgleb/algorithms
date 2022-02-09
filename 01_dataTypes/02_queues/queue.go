package queue

type Queue interface {
	Enqueue(string)
	Dequeue() (string, error)
	IsEmpty() bool
	Size() int
	Iterate() string
}

func Enqueue(q Queue, val string) {
	q.Enqueue(val)
}

func Dequeue(q Queue) (string, error) {
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
