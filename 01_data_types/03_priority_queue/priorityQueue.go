package priorityQueue

type PriorityQueue interface {
	Insert(int)
	DelMax() (int, error)
	IsEmpty() bool
	Max() int
	Size() int
}

func Insert(p PriorityQueue, val int) {
	p.Insert(val)
}

func DelMax(p PriorityQueue) (int, error) {
	return p.DelMax()
}

func IsEmpty(p PriorityQueue) bool {
	return p.IsEmpty()
}

func Max(p PriorityQueue) int {
	return p.Max()
}

func Size(p PriorityQueue) int {
	return p.Size()
}
