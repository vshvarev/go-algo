package validate_subsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	i := 0

	for _, v := range array {
		if v == sequence[i] {
			i++
		}
		if i == len(sequence) {
			return true
		}
	}

	return false
}
