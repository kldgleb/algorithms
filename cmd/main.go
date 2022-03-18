package main

import (
	priorityQueue "algorithms/01_dataTypes/03_priorityQueue"
	"fmt"
)

func main() {
	heapPriorityQueue := priorityQueue.NewHeap()
	testPriorityQueue(heapPriorityQueue)
	testHeapSort()
}

func testPriorityQueue(h priorityQueue.PriorityQueue) {
	fmt.Println("\n \n Testing PriorityQueue...")

	priorityQueue.Insert(h, 3)
	fmt.Println("Add value: ", 3)

	priorityQueue.Insert(h, 4)
	fmt.Println("Add value: ", 4)

	priorityQueue.Insert(h, 1)
	fmt.Println("Add value: ", 1)

	max, _ := h.Max()
	fmt.Println("Max value: ", max)
	max, err := h.DelMax()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted value: ", max)
	}
}

func testHeapSort() {
	fmt.Println("\n \n Testing Heapsort...")

	array := []int{7, 2, 3, 4, 5, 6, 1}
	h := priorityQueue.Sort(array)
	fmt.Println(h)
}
