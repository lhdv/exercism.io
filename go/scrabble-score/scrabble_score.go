package scrabble

import "strings"

// Score computes a string score
func Score(s string) int {

	score := 0

	if len(s) <= 0 {
		return 0
	}

	for _, c := range s {
		score += getLetterScore(string(c))
	}

	return score

}

func getLetterScore(c string) int {

	c = strings.ToUpper(c)
	s := 0

	switch {
	case strings.Contains("AEIOULNRST", c):
		s = 1
	case strings.Contains("DG", c):
		s = 2
	case strings.Contains("BCMP", c):
		s = 3
	case strings.Contains("FHVWY", c):
		s = 4
	case strings.Contains("K", c):
		s = 5
	case strings.Contains("JX", c):
		s = 8
	case strings.Contains("QZ", c):
		s = 10
	}

	return s
}
