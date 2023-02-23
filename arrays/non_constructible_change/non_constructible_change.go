package non_constructible_change

import (
	"sort"
)

func NonConstructibleChange(coins []int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})

	change := 0

	for _, v := range coins {
		if v > change+1 {
			return change + 1
		}

		change += v
	}

	return change + 1
}
