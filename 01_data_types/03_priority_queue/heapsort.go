package priorityQueue

import "fmt"

func Sort(slice []int) *heap {
	empty := []int{0}
	empty = append(empty, slice...)
	heap := &heap{empty, len(empty) - 1}
	fmt.Println(heap)

	for k := heap.n / 2; k >= 1; k-- {
		heap.sink(k)
		fmt.Println("sink: ", k)
		for heap.n > 1 {
			heap.exch(1, heap.n)
			heap.n--
			heap.sink(1)
		}
	}
	return heap
}
