package pascal

import (
	"log"
)

// Triangle build a Pascal Triangle for a given int
func Triangle(n int) [][]int {

	var result [][]int

	if n < 1 {
		return [][]int{}
	}

	result = append(result, []int{1})

	log.Printf("%v\n", result)

	for i := 1; i < n; i++ {
		var innerResult []int
		for j := 1; j <= n; i++ {
			result[i] = append(result[i], i)
		}
		result = append(result, innerResult)
	}

	return result
}
