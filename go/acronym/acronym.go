// Package acronym build acronyms from a series of words
package acronym

import (
	"regexp"
	"strings"
)

func extractWords(s string) [][]string {
	exp := `([a-zA-Z])[a-zA-Z]+`
	re := regexp.MustCompile(exp)
	return re.FindAllStringSubmatch(s, -1)
}

// Abbreviate do it's job
func Abbreviate(s string) string {

	var result string

	words := extractWords(s)
	for _, w := range words {
		result += strings.ToUpper(w[1])
	}

	return result
}
