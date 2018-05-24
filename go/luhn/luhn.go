package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Valid check luhn's algorithm
func Valid(s string) bool {

	s = stripAllBlanks(s)

	re := regexp.MustCompile(`[^\d\s]`)
	if re.MatchString(s) || len(s) <= 1 {
		return false
	}

	split := strings.Split(s, "")

	for i := len(split) - 2; i >= 0; i = i - 2 {
		split[i] = doubleDigit(split[i])
	}

	valid := checkValidDigits(split)

	return valid
}

func stripAllBlanks(s string) string {
	re := regexp.MustCompile(`\s`)
	return re.ReplaceAllString(s, "")
}

func doubleDigit(s string) string {

	n, _ := strconv.Atoi(s)
	n = n * 2
	if n > 9 {
		n = n - 9
	}

	r := strconv.Itoa(n)

	return r
}

func checkValidDigits(digits []string) bool {

	sum := 0
	n := 0

	for _, d := range digits {
		n, _ = strconv.Atoi(d)
		sum += n
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
