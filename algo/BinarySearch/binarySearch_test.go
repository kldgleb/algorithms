package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		arr    []int
		target int
		res    bool
	}{
		{
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 5,
			res:    true,
		},
		{
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 0,
			res:    false,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 5,
			res:    true,
		},
	}

	for i, test := range testTable {
		test := test
		t.Run(fmt.Sprintf("Case #%d, target - %d", i, test.target), func(t *testing.T) {
			t.Parallel()
			res := BinarySearch(test.arr, test.target)
			assert.Equal(t, test.res, res)
		})
	}
}

func TestSearchInsert(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		arr    []int
		target int
		res    int
	}{
		{
			arr:    []int{1, 3, 5, 6},
			target: 5,
			res:    2,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 2,
			res:    1,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 7,
			res:    4,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 0,
			res:    0,
		},
	}

	for i, test := range testTable {
		test := test
		t.Run(fmt.Sprintf("Case #%d, target - %d", i, test.target), func(t *testing.T) {
			t.Parallel()
			res := SearchInsert(test.arr, test.target)
			assert.Equal(t, test.res, res)
		})
	}
}
