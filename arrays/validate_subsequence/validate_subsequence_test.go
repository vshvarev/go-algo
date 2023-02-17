package validate_subsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type isValidSubsequenceTest struct {
	arr, seq []int
	expected bool
}

var isValidSubsequenceTests = []isValidSubsequenceTest{
	{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}, true},
	{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10}, true},
	{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{25}, true},
	{[]int{1, 1, 1, 1}, []int{1, 1, 1}, true},
	{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{26}, false},
	{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10, 5}, false},
	{[]int{1, 1, 6, 1}, []int{1, 1, 1, 6}, false},
}

func TestIsValidSubsequence(t *testing.T) {
	for _, test := range isValidSubsequenceTests {
		actual := IsValidSubsequence(test.arr, test.seq)
		require.Equal(t, actual, test.expected)
	}
}
