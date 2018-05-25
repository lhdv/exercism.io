package grains

import (
	"errors"
)

// Square calculate how many wheat grains there are in a 8x8 chess board
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("invalid number")
	}

	return (1 << (uint64(n) - 1)), nil
}

// Total sums all the wheat grains a 8x8 chess board has
func Total() uint64 {
	return (uint64)((1 << 64) - 1)
}
