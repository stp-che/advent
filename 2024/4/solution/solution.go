package solution

import (
	"log"
	"os"
	"strings"
)

var directions = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

const xmas = "XMAS"

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	input := s.readData(inputPath)
	sum := 0
	for i, row := range input {
		for j, c := range row {
			if c == 'X' {
				sum += s.search(input, i, j)
			}
		}
	}

	return sum
}

func (s *Solution) Run1(inputPath string) int {
	table := s.readData(inputPath)
	height := len(table)
	width := len(table[0])
	sum := 0

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if s.isXMasCross(table, i, j) {
				sum++
			}
		}
	}

	return sum
}

func (s *Solution) search(table []string, row, col int) int {
	res := 0
	height := len(table)
	width := len(table[0])
	searchLen := len(xmas)

	for _, d := range directions {
		i := row + d[0]*(searchLen-1)
		j := col + d[1]*(searchLen-1)
		if i < 0 || i >= height || j < 0 || j >= width {
			continue
		}

		n := 1
		i, j = row+d[0], col+d[1]
		for n < searchLen && table[i][j] == xmas[n] {
			n++
			i += d[0]
			j += d[1]
		}
		if n == searchLen {
			res++
		}
	}

	return res
}

func (s *Solution) isXMasCross(table []string, row, col int) bool {
	topLeft := table[row-1][col-1]
	bottomRight := table[row+1][col+1]
	topRight := table[row-1][col+1]
	bottomLeft := table[row+1][col-1]
	return table[row][col] == 'A' &&
		(topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M') &&
		(topRight == 'M' && bottomLeft == 'S' || topRight == 'S' && bottomLeft == 'M')
}

func (s *Solution) readData(inputPath string) []string {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
