package robotname

import (
	"fmt"
	"math/rand"
)

// Robot type
type Robot struct {
	name string
}

var robotNamePool = make(map[string]bool)

// Name show robot's name
func (r *Robot) Name() string {

	if len(r.name) <= 0 {
		r.name = generateName()
		for robotNamePool[r.name] {
			r.name = generateName()
		}
		robotNamePool[r.name] = true
	}
	return r.name
}

// Reset robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	l1 := rand.Intn(25) + 'A'
	l2 := rand.Intn(25) + 'A'
	d := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", string(l1), string(l2), d)
}
