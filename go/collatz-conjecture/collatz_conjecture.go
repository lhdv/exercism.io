package collatzconjecture

import (
	"errors"
)

// CollatzConjecture will return how many steps is required for a given integer number return 1
func CollatzConjecture(n int) (int, error) {

	var steps int

	if n <= 0 {
		return 0, errors.New("Invalid input data")
	}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		steps++
	}

	return steps, nil
}
