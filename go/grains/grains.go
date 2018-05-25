package grains

import (
	"errors"
	"math"
)

// Square calculate how many wheat grains there are in a 8x8 chess board
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("invalid number")
	}

	square := math.Pow(2, float64(n-1))

	return uint64(square), nil
}

// Total sums all the wheat grains a 8x8 chess board has
func Total() uint64 {
	var r uint64
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		r += s
	}

	return r
}
