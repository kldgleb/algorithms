package priorityQueue

import (
	"errors"
	"fmt"
)

type Heap struct {
	Slice []int
	N     int
}

func NewHeap() *Heap {
	return &Heap{[]int{0}, 0}
}

func (h *Heap) Insert(val int) {
	h.Slice = append(h.Slice, val)
	h.N++
	h.Swim(h.N)
}

func (h *Heap) DelMax() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	buffer := h.Slice[1]
	h.Exch(1, h.N)
	h.Slice = h.Slice[:h.N]
	h.N--
	h.Sink(1)
	return buffer, nil
}

func (h Heap) IsEmpty() bool {
	return h.N == 0
}

func (h Heap) Max() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return h.Slice[1], nil
}

func (h *Heap) Size() int {
	return h.N
}

func (h *Heap) Sink(key int) {
	for 2*key <= h.N {
		j := 2 * key
		if j < h.N && h.Slice[j] < h.Slice[j+1] {
			j++
		}
		if key > j {
			break
		}
		fmt.Printf("Sink change: %d with %d \n", h.Slice[key], h.Slice[j])
		h.Exch(key, j)
		key = j
	}
}

func (h *Heap) Swim(key int) {
	for key > 1 && h.Slice[key/2] < h.Slice[key] {
		h.Exch(key, key/2)
		key = key / 2
	}
}

func (h *Heap) Exch(a, b int) {
	h.Slice[b], h.Slice[a] = h.Slice[a], h.Slice[b]
}

func parent(child int) int {
	return child*2 + 1
}

func leftChild(parent int) int {
	return parent/2 + 1
}

func rightChild(parent int) int {
	return parent/2 + 2
}
