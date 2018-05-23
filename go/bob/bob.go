package bob

import (
	"regexp"
	"strings"
)

func isQuestioning(r string) bool {
	re := regexp.MustCompile(`\?$`)
	return re.MatchString(r)
}

func isShouting(r string) bool {
	r = justLetters(r)
	re := regexp.MustCompile(`^[A-Z]+$`)
	return re.MatchString(r)
}

func isSayingNothing(r string) bool {
	re := regexp.MustCompile(`^\s*$`)
	return re.MatchString(r)
}

func justLetters(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z]`)
	return re.ReplaceAllString(s, "")
}

// Hey will return what Bob would say
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	switch {
	case isQuestioning(remark) && isShouting(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestioning(remark):
		return "Sure."
	case isShouting(remark):
		return "Whoa, chill out!"
	case isSayingNothing(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
