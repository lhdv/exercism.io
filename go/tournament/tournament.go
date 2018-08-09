package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// team defines a team entry in score table
type team struct {
	Name    string
	Matches int
	Wins    int
	Draws   int
	Loses   int
	Points  int
}

func (t team) String() string {
	return fmt.Sprintf("%-31v| %2d | %2d | %2d | %2d | %2d", t.Name, t.Matches, t.Wins, t.Draws, t.Loses, t.Points)
}

// Tally build a score table from a soccer competition
func Tally(r io.Reader, w io.Writer) error {

	teams := make(map[string]team, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()

		// Ignore empty lines and lines beginning with '#'
		if len(s) > 0 && s[:1] != "#" {
			err := buildTeam(s, teams)
			if err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return errors.New("Error reading input data")
	}

	printOutput(teams, w)

	return nil
}

func buildTeam(input string, t map[string]team) error {

	var ta, tb team

	input = strings.TrimSpace(input)
	input = strings.TrimSuffix(input, "\n")
	s := strings.Split(input, ";")

	if len(s) == 3 {
		ta = t[s[0]]
		tb = t[s[1]]
		result := s[2]

		ta.Name = s[0]
		ta.Matches++
		tb.Name = s[1]
		tb.Matches++

		switch {
		case result == "win":
			ta.Points += 3
			ta.Wins++
			tb.Loses++
		case result == "loss":
			ta.Loses++
			tb.Points += 3
			tb.Wins++
		case result == "draw":
			ta.Draws++
			ta.Points++
			tb.Draws++
			tb.Points++
		default:
			return errors.New("Invalid Result Data")
		}

		t[s[0]] = ta
		t[s[1]] = tb
	} else {
		return errors.New("Invalid Input Data")
	}

	return nil
}

func printOutput(teams map[string]team, w io.Writer) {

	table := make([]team, 0)

	for _, t := range teams {
		table = append(table, t)
	}

	sort.SliceStable(table, func(i, j int) bool { return table[i].Name < table[j].Name })
	sort.SliceStable(table, func(i, j int) bool { return table[i].Points > table[j].Points })

	write := bufio.NewWriter(w)
	fmt.Fprintln(write, fmt.Sprintf("%-31v| %2v | %2v | %2v | %2v | %2v", "Team", "MP", "W", "D", "L", "P"))
	for _, t := range table {
		fmt.Fprintln(write, t)
	}
	write.Flush()

}
