package twelve

import (
	"bytes"
	"fmt"
)

var (
	gifts = []string{
		"a Partridge in a Pear Tree.",
		"two Turtle Doves, ",
		"three French Hens, ",
		"four Calling Birds, ",
		"five Gold Rings, ",
		"six Geese-a-Laying, ",
		"seven Swans-a-Swimming, ",
		"eight Maids-a-Milking, ",
		"nine Ladies Dancing, ",
		"ten Lords-a-Leaping, ",
		"eleven Pipers Piping, ",
		"twelve Drummers Drumming, "}

	ordinalNumber = []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth"}
)

type ordinal int

func (o ordinal) String() string {
	return ordinalNumber[o-1]
}

type lyricVerse struct {
	day  int
	gift string
}

func (lv lyricVerse) String() string {

	var b bytes.Buffer

	fmt.Fprintf(&b, "On the %s day of Christmas my true love gave to me, ", ordinal(lv.day))

	for i := lv.day; i >= 1; i-- {
		if i <= 1 && lv.day > 1 {
			fmt.Fprintf(&b, "%s", "and ")
		}
		fmt.Fprintf(&b, "%s", gifts[i-1])
	}

	return b.String()
}

// Song print the twelve days song
func Song() string {

	var song bytes.Buffer

	for i := 1; i <= 12; i++ {
		fmt.Fprintf(&song, "%s\n", Verse(i))
	}

	return song.String()
}

// Verse return a song's verse based on its number
func Verse(n int) string {

	if n < 1 || n > 12 {
		return "Wrong verse number"
	}

	return lyricVerse{day: n}.String()
}
