package sort

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		data []int
		res  []int
	}{
		// {
		// 	data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 	res:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// },
		{
			data: []int{6, 8, 5, 10, 9, 2, 0, 7, 1, 4, 3},
			res:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		// {
		// 	data: []int{2, 3, 0, 1, 7, 8, 6, 4, 5, 10, 9},
		// 	res:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// },
	}

	for i, test := range testTable {
		test := test
		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
			t.Parallel()
			s := NewHeapSort()
			Sort(s, test.data)
			fmt.Println(test.data)
			fmt.Println(s.heap.Slice)
			ok := reflect.DeepEqual(test.data, test.res)
			assert.Equal(t, true, ok)
		})
	}

}
