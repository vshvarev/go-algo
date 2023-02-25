package sorted_squared_array

func SortedSquaredArray(array []int) []int {
	l := len(array)
	res := make([]int, l)

	start := 0
	end := l - 1

	for i := range array {
		if abs(array[start]) > abs(array[end]) {
			res[l-i-1] = array[start] * array[start]
			start++
		} else {
			res[l-i-1] = array[end] * array[end]
			end--
		}
	}

	return res
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return n * -1
}
