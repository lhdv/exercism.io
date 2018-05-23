package diffsquares

import "math"

// SquareOfSums first sum all numbers then calculate the square
func SquareOfSums(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		ret += i
	}

	return square(ret)
}

// SumOfSquares first calculate the square of number then sum all of them
func SumOfSquares(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		ret += square(i)
	}

	return ret
}

// Difference calculate SquareOfSums - SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func square(n int) int {
	return int(math.Pow(float64(n), 2))
}
