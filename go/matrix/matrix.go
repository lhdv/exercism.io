package matrix

import (
	"encoding/csv"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Matrix type struct
type Matrix [][]int

// New receives a string matrix delimited by new-line caracter and returns
// a slice with rows and columns content
func New(input string) (Matrix, error) {

	if hasEmptyRows(input) {
		return nil, errors.New("Empty rows was found")
	}

	recReader := csv.NewReader(strings.NewReader(input))
	recReader.Comma = ' '
	recReader.TrimLeadingSpace = true

	records, err := recReader.ReadAll()
	if err != nil {
		return nil, errors.New("Error reading input data")
	}

	matrix := make(Matrix, 0)

	for _, record := range records {
		numbers, err := convertSliceStringToInt(record)
		if err != nil {
			return nil, err
		}
		matrix = append(matrix, numbers)
	}

	return matrix, nil
}

// Rows returns a slice of int from matrix data
func (m Matrix) Rows() [][]int {
	var output Matrix

	for _, row := range m {
		c := make([]int, 0)
		for _, col := range row {
			c = append(c, col)
		}

		output = append(output, c)
	}
	return output
}

// Cols returns a slice of int from matrix data
func (m Matrix) Cols() [][]int {
	var output Matrix

	for i := 0; i < len(m[0]); i++ {
		c := make([]int, 0)
		for j := 0; j < len(m); j++ {
			c = append(c, m[j][i])
		}

		output = append(output, c)
	}

	return output
}

// Set a given value on an specific matrix position
func (m *Matrix) Set(a, b, c int) bool {

	for kr, row := range *m {
		for kc := range row {
			if kr == a && kc == b {
				(*m)[kr][kc] = c
				return true
			}
		}
	}

	return false
}

func hasEmptyRows(input string) bool {
	if strings.HasPrefix(input, "\n") || strings.HasSuffix(input, "\n") {
		return true
	}

	if inputInvalid, _ := regexp.MatchString("\\n\\n", input); inputInvalid {
		return true
	}

	return false
}

func convertSliceStringToInt(input []string) ([]int, error) {
	ret := make([]int, 0)

	for _, v := range input {
		n, error := strconv.Atoi(v)
		if error != nil {
			return nil, errors.New("Invalid int data)")
		}
		ret = append(ret, n)
	}

	return ret, nil
}