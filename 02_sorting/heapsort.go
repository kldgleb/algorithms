package sort

import priorityQueue "github.com/kldgleb/algorithms/01_dataTypes/03_priorityQueue"

type HeapSort struct {
	heap *priorityQueue.Heap
}

func NewHeapSort() *HeapSort {
	heap := &priorityQueue.Heap{
		Slice: []int{0},
		N:     0,
	}
	return &HeapSort{
		heap: heap,
	}
}

func (s *HeapSort) Sort(slice []int) {
	s.heap.Slice = append(s.heap.Slice, slice...)
	s.heap.N += len(slice)
	for k := s.heap.N / 2; k >= 1; k-- {
		s.heap.Sink(k)
		for s.heap.N > 1 {
			s.heap.Exch(1, s.heap.N)
			s.heap.N--
			s.heap.Sink(1)
		}
	}
}
