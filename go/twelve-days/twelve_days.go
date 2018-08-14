package twelve

import (
	"fmt"
)

type lyric struct {
	ordinalDay string
	gift       string
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

	var verse string
	var gifts string

	if n < 1 || n > 12 {
		return ""
	}

	lyrics := make(map[int]lyric)
	lyrics[1] = lyric{ordinalDay: "first", gift: "a Partridge in a Pear Tree"}
	lyrics[2] = lyric{ordinalDay: "second", gift: "two Turtle Doves"}
	lyrics[3] = lyric{ordinalDay: "third", gift: "three French Hens"}
	lyrics[4] = lyric{ordinalDay: "fourth", gift: "four Calling Birds"}
	lyrics[5] = lyric{ordinalDay: "fifth", gift: "five Gold Rings"}
	lyrics[6] = lyric{ordinalDay: "sixth", gift: "six Geese-a-Laying"}
	lyrics[7] = lyric{ordinalDay: "seventh", gift: "seven Swans-a-Swimming"}
	lyrics[8] = lyric{ordinalDay: "eighth", gift: "eight Maids-a-Milking"}
	lyrics[9] = lyric{ordinalDay: "ninth", gift: "nine Ladies Dancing"}
	lyrics[10] = lyric{ordinalDay: "tenth", gift: "ten Lords-a-Leaping"}
	lyrics[11] = lyric{ordinalDay: "eleventh", gift: "eleven Pipers Piping"}
	lyrics[12] = lyric{ordinalDay: "twelfth", gift: "twelve Drummers Drumming"}

	for i := n; i >= 1; i-- {
		if i > 1 {
			gifts += lyrics[i].gift + ", "
		} else {
			if n > 1 {
				gifts += "and " + lyrics[i].gift + "."
			} else {
				gifts += lyrics[i].gift + "."
			}
		}

	}

	verse = fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s", lyrics[n].ordinalDay, gifts)

	return verse

}
