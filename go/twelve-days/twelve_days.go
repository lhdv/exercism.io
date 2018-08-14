package twelve

import (
	"fmt"
)

type ordinal int

func (o ordinal) String() string {
	switch {
	case o == 1:
		return "first"
	case o == 2:
		return "second"
	case o == 3:
		return "third"
	case o == 4:
		return "fourth"
	case o == 5:
		return "fifth"
	case o == 6:
		return "sixth"
	case o == 7:
		return "seventh"
	case o == 8:
		return "eighth"
	case o == 9:
		return "ninth"
	case o == 10:
		return "tenth"
	case o == 11:
		return "eleventh"
	case o == 12:
		return "twelfth"
	default:
		return ""
	}
}

type lyricVerse struct {
	day  int
	gift string
}

func (lv lyricVerse) String() string {

	gifts := make(map[int]string)
	gifts[1] = "a Partridge in a Pear Tree"
	gifts[2] = "two Turtle Doves"
	gifts[3] = "three French Hens"
	gifts[4] = "four Calling Birds"
	gifts[5] = "five Gold Rings"
	gifts[6] = "six Geese-a-Laying"
	gifts[7] = "seven Swans-a-Swimming"
	gifts[8] = "eight Maids-a-Milking"
	gifts[9] = "nine Ladies Dancing"
	gifts[10] = "ten Lords-a-Leaping"
	gifts[11] = "eleven Pipers Piping"
	gifts[12] = "twelve Drummers Drumming"

	for i := lv.day; i >= 1; i-- {
		if i > 1 {
			lv.gift += gifts[i] + ", "
		} else {
			if lv.day > 1 {
				lv.gift += "and " + gifts[i] + "."
			} else {
				lv.gift += gifts[i] + "."
			}
		}
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s", ordinal(lv.day), lv.gift)
}

// Song print the twelve days song
func Song() string {

	var song string

	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}

	return song
}

// Verse return a song's verse based on its number
func Verse(n int) string {

	if n < 1 || n > 12 {
		return "Wrong verse number"
	}

	return lyricVerse{day: n}.String()
}
