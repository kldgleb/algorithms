package twopointers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstProblem(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		time  int
		books []int
		res   int
	}{
		{
			time:  5,
			books: []int{3, 1, 2, 1},
			res:   3,
		},
		{
			time:  3,
			books: []int{2, 2, 3},
			res:   1,
		},
	}

	for i, test := range testTable {
		test := test
		t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
			t.Parallel()
			res := FirstProblem(test.time, test.books)
			assert.Equal(t, test.res, res)
		})
	}

}
