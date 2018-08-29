package pascal

// Triangle build a Pascal Triangle for a given int
func Triangle(n int) [][]int {

	var result [][]int

	if n < 1 {
		return [][]int{}
	}

	result = append(result, []int{1})
	if n == 1 {
		return result
	}

	result = append(result, []int{1, 1})
	if n == 2 {
		return result
	}

	for i := 2; i < n; i++ {
		row := make([]int, 0)
		row = append(row, 1)
		for j := 1; j < i; j++ {
			n := result[i-1][j-1] + result[i-1][j]
			row = append(row, n)
		}
		row = append(row, 1)
		result = append(result, row)
	}

	return result
}
