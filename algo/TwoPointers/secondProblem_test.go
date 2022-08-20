package twopointers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondProblem(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		students []int
		res      int
	}{
		{
			students: []int{1, 10, 17, 12, 15, 2},
			res:      3,
		},
		{
			students: []int{1337, 1337, 1337, 1337, 1337, 1337, 1337, 1337, 1337, 1337},
			res:      10,
		},
		{
			students: []int{1, 1000, 10000, 10, 100, 1000000000},
			res:      1,
		},
	}
	for i, test := range testTable {
		test := test
		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
			t.Parallel()
			res := SecondProblem(test.students)
			assert.Equal(t, test.res, res)
		})
	}
}

// students 1 10 17 12 15 2
// res - 3

// 1337 1337 1337 1337 1337 1337 1337 1337 1337 1337
// res - 10

// 1 1000 10000 10 100 1000000000
// res - 1
