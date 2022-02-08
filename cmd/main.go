package main

import (
	stack "algorithms/01_data_types/01_stack"
	queue "algorithms/01_data_types/02_queues"
	priorityQueue "algorithms/01_data_types/03_priority_queue"
	"fmt"
	"log"
)

func main() {
	arrayStack := stack.NewArrayStack()
	testStack(arrayStack)
	linkedListStack := stack.NewLinkedListStack()
	testStack(linkedListStack)

	arrayQueue := queue.NewArrayQueue()
	testQueue(arrayQueue)
	linkedListQueue := queue.NewLinkedListQueue()
	testQueue(linkedListQueue)

	heapPriorityQueue := priorityQueue.NewHeap()
	testPriorityQueue(heapPriorityQueue)
	testHeapSort()
}

func testStack(s stack.Stack) {
	fmt.Println("\n \n Testing Stack...")
	stack.Size(s)
	stack.Push(s, "val1")
	stack.Push(s, "val2")
	top, err := stack.Top(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Top value: ", top)
	deletedVal, err := stack.Pop(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted value: ", deletedVal)
	stack.Print(s)
}

func testQueue(q queue.Queue) {
	fmt.Println("\n \n Testing Queue...")

	queue.Enqueue(q, "val1")
	queue.Enqueue(q, "val2")
	queue.Enqueue(q, "val3")
	deletedVal, err := queue.Dequeue(q)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted val: ", deletedVal)
	fmt.Println("Size: ", queue.Size(q))
	queue.Iterate(q)
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
