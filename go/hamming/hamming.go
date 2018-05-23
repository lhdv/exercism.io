package hamming

import (
	"bytes"
	"errors"
	"strings"
)

// Distance do it's job
func Distance(a, b string) (int, error) {

	var d int

	if len(a) != len(b) {
		return -1, errors.New("different strands size")
	}

	a = strings.ToUpper(a)
	b = strings.ToUpper(b)

	runesFromA := bytes.Runes(bytes.NewBufferString(a).Bytes())
	runesFromB := bytes.Runes(bytes.NewBufferString(b).Bytes())

	for i, b := range runesFromA {
		if b != runesFromB[i] {
			d++
		}
	}

	return d, nil
}
