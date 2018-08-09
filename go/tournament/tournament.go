package tournament

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
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

	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'
	csvReader.Comment = '#'

	records, err := csvReader.ReadAll()
	if err != nil {
		return errors.New("Error reading input data")
	}

	for _, record := range records {
		err := buildTeam(record, teams)
		if err != nil {
			return err
		}
	}

	printOutput(teams, w)

	return nil
}

func buildTeam(input []string, t map[string]team) error {

	var teamA, teamB team

	if len(input) != 3 {
		return errors.New("Invalid Input Data")
	}

	teamA = t[input[0]]
	teamB = t[input[1]]
	result := input[2]

	teamA.Name = input[0]
	teamA.Matches++
	teamB.Name = input[1]
	teamB.Matches++

	switch {
	case result == "win":
		teamA.Points += 3
		teamA.Wins++
		teamB.Loses++
	case result == "loss":
		teamA.Loses++
		teamB.Points += 3
		teamB.Wins++
	case result == "draw":
		teamA.Draws++
		teamA.Points++
		teamB.Draws++
		teamB.Points++
	default:
		return errors.New("Invalid Result Data")
	}

	t[input[0]] = teamA
	t[input[1]] = teamB

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
