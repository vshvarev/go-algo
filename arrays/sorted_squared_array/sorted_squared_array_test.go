package sorted_squared_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type sortedSquaredArrayTest struct {
	arr      []int
	expected []int
}

var sortedSquaredArrayTests = []sortedSquaredArrayTest{
	{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 4, 9, 16, 25, 36, 49}},
	{[]int{-4, 1, 2}, []int{1, 4, 16}},
}

func TestSortedSquaredArray(t *testing.T) {
	for _, test := range sortedSquaredArrayTests {
		actual := SortedSquaredArray(test.arr)
		require.Equal(t, actual, test.expected)
	}
}
