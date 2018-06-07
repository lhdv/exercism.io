package robotname

import (
	"math/rand"
	"strconv"
)

// Robot type
type Robot struct {
	name string
}

// Name show robot's name
func (r *Robot) Name() string {

	if len(r.name) <= 0 {

		l1 := rand.Intn(25) + 65
		l2 := rand.Intn(25) + 65
		d1 := rand.Intn(9)
		d2 := rand.Intn(9)
		d3 := rand.Intn(9)
		s := string(l1) + string(l2) + strconv.Itoa(d1) + strconv.Itoa(d2) + strconv.Itoa(d3)

		r.name = s
	}
	return r.name
}

// Reset robot's name
func (r *Robot) Reset() {
	r.name = ""
}
