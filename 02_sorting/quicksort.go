package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	fmt.Println("shuffled: ", a)
	quicksort(a)
	fmt.Println(a)
}

func quicksort(a []int) {
	sort(a, 0, len(a)-1)
}

func sort(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	sort(a, lo, j-1)
	sort(a, j+1, hi)
}
func partition(a []int, lo, hi int) int {
	j := hi
	i := lo + 1
	for {
		// find item on left to swap
		for a[i] < a[lo] {
			i++
			if i == hi {
				break
			}
		}
		//find item on right to swap
		for a[lo] < a[j] {
			j--
			if j == lo {
				break
			}
		}

		//check if pointer cross
		if i >= j {
			break
		}
		exch(a, i, j)
	}

	exch(a, lo, j)
	return j // the index of item, separating the array
}

func exch(a []int, i, j int) {
	buffer := a[i]
	a[i] = a[j]
	a[j] = buffer
}
