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

	h.Insert(1)
	fmt.Println("Add value: ", 1)

	h.Insert(3)
	fmt.Println("Add value: ", 3)

	h.Insert(2)
	fmt.Println("Add value: ", 2)

	fmt.Println("Max value: ", h.Max())
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
