package two_number_sum

func TwoNumberSum(arr []int, target int) []int {
	m := map[int]struct{}{}
	for _, currentNum := range arr {
		numNeeded := target - currentNum

		if _, ok := m[numNeeded]; ok {
			return []int{numNeeded, currentNum}
		}

		m[currentNum] = struct{}{}
	}

	return []int{}
}
