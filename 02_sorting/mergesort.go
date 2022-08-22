package sort

type MergeSort struct {
}

func NewMergeSort() *MergeSort {
	return &MergeSort{}
}

func (m *MergeSort) Sort(data []int) {
	m.sort(data, []int{}, 0, len(data)-1)
}

func (m *MergeSort) sort(a, sub []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	m.sort(a, sub, lo, mid)
	m.sort(a, sub, mid+1, hi)
	m.merge(a, lo, mid, hi)
}

func (m *MergeSort) merge(a []int, lo, mid, hi int) {
	sub := make([]int, hi+1)
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
