package main

import "fmt"

//often exit with error index out of range [10] with length 10, line 35
// idk why

func main() {
	a := []int{8, 4, 10, 9, 7, 2, 6, 3, 5, 1}
	empty := []int{}
	sort(a, empty, 0, len(a)-1)
	fmt.Println(a)
}

func sort(a, sub []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, sub, lo, mid)
	sort(a, sub, mid+1, hi)
	merge(a, sub, lo, mid, hi)
}

func merge(a, sub []int, lo, mid, hi int) {
	sub = make([]int, hi+1)
	for k := lo; k <= hi; k++ {
		sub[k] = a[k]
	}
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = sub[j]
			j++
		} else if j > hi {
			a[k] = sub[i]
			i++
		} else if sub[j] > sub[i] {
			a[k] = sub[i]
			i++
		} else {
			a[k] = sub[j]
			j++
		}
	}
}
