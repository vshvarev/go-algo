package non_constructible_change

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type NonConstructibleChangeTest struct {
	coins    []int
	expected int
}

var NonConstructibleChangeTests = []NonConstructibleChangeTest{
	{[]int{5, 7, 1, 1, 2, 3, 22}, 20},
	{[]int{1, 1, 1, 1, 1}, 6},
	{[]int{87}, 1},
	{[]int{}, 1},
}

func TestNonConstructibleChange(t *testing.T) {
	for _, test := range NonConstructibleChangeTests {
		actual := NonConstructibleChange(test.coins)
		require.Equal(t, actual, test.expected)
	}
}
