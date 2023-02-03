package two_number_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	actual := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	expected := []int{-1, 11}

	require.ElementsMatch(t, actual, expected)
}
