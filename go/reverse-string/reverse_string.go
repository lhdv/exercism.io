package reverse

import (
	"bytes"
)

// String reverse a given string
func String(s string) string {

	var ret string

	buffer := bytes.NewBufferString(s)
	sInRunes := bytes.Runes(buffer.Bytes())

	for i := len(sInRunes) - 1; i >= 0; i-- {
		ret += string(sInRunes[i])
	}

	return ret
}
