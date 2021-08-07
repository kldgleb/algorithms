package main

import (
	"errors"
	"fmt"
)

type heap struct {
	slice []int
	n     int
}

func main() {
	empty := []int{0}
	heap := &heap{empty, 0}

	heap.insert(1)
	heap.insert(3)
	heap.insert(2)

	max, err := heap.delMax()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(max)
	}

	result := heap.slice[1 : heap.n+1]
	fmt.Println(result)
}

func (h *heap) insert(val int) {
	h.slice = append(h.slice, val)
	h.n++
	h.swim(h.n)
}

func (h *heap) delMax() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("empty queue")
	}
	buffer := h.slice[1]
	h.exch(1, h.n)
	h.n--
	h.slice = h.slice[:h.n+1]
	h.sink(1)
	return buffer, nil
}

func (h *heap) sink(key int) {
	for 2*key <= h.n {
		j := 2 * key
		if j < h.n && h.slice[j] < h.slice[j+1] {
			j++
		}
		if h.slice[j] < h.slice[key] {
			break
		}
		key = j
		h.exch(key, j)
	}
}

func (h *heap) swim(key int) {
	for key > 1 && h.slice[key/2] < h.slice[key] {
		h.exch(key, key/2)
		key = key / 2
	}
}

func (h *heap) exch(a, b int) {
	buffer := h.slice[b]
	h.slice[b] = h.slice[a]
	h.slice[a] = buffer
}

func (h *heap) isEmpty() bool {
	return h.n == 0
}
