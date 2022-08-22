package sort

type Sorter interface {
	Sort([]int)
}

func Sort(s Sorter, data []int) {
	s.Sort(data)
}
