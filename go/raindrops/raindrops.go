package raindrops

import "strconv"

// Convert do it's job
func Convert(n int) string {

	var r string

	if n%3 == 0 {
		r += "Pling"
	}

	if n%5 == 0 {
		r += "Plang"
	}

	if n%7 == 0 {
		r += "Plong"
	}

	if len(r) <= 0 {
		r += strconv.Itoa(n)
	}

	return r
}
