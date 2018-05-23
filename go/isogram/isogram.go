package isogram

import "strings"

// IsIsogram check if a text is isogram
func IsIsogram(s string) bool {

	var text string

	for _, c := range s {
		if strings.Contains(text, string(c)) {
			return false
		}

		text += string(c)
	}

	return true
}
