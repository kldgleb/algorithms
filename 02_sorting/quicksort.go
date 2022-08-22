package sort

type QuickSort struct {
}

func NewQuickSort() *QuickSort {
	return &QuickSort{}
}

func (q *QuickSort) Sort(a []int) {
	q.sort(a, 0, len(a)-1)
}

func (q *QuickSort) sort(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	j := q.partition(a, lo, hi)
	q.sort(a, lo, j-1)
	q.sort(a, j+1, hi)
}

func (q *QuickSort) partition(a []int, lo, hi int) int {
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
		q.exch(a, i, j)
	}

	q.exch(a, lo, j)
	return j // the index of item, separating the array
}

func (q *QuickSort) exch(a []int, i, j int) {
	buffer := a[i]
	a[i] = a[j]
	a[j] = buffer
}
