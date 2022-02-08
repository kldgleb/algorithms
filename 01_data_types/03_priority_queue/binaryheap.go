package priorityQueue

import (
	"errors"
)

type heap struct {
	slice []int
	n     int
}

func NewHeap() *heap {
	return &heap{[]int{0}, 0}
}

func (h *heap) Insert(val int) {
	h.slice = append(h.slice, val)
	h.n++
	h.swim(h.n)
}

func (h *heap) DelMax() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	buffer := h.slice[1]
	h.exch(1, h.n)
	h.n--
	h.slice = h.slice[:h.n+1]
	h.sink(1)
	return buffer, nil
}

func (h heap) IsEmpty() bool {
	return h.n == 0
}

func (h heap) Max() int {
	return h.slice[1]
}

func (h *heap) Size() int {
	return h.n
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
