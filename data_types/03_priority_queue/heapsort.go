package main

import "fmt"

type heap struct {
	slice []int
	n     int
}

func main() {
	array := []int{7, 2, 3, 4, 5, 6, 1}
	h := sort(array)
	fmt.Println(h.slice)
}

func sort(slice []int) *heap {

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

func (h *heap) exch(a, b int) {
	fmt.Println("exch: ", h.slice[a], "and", h.slice[b])
	buffer := h.slice[b]
	h.slice[b] = h.slice[a]
	h.slice[a] = buffer
}
