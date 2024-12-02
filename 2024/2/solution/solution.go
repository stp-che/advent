package solution

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	reports := s.readData(inputPath)

	sum := 0
	for _, r := range reports {
		if s.isSafe(r) {
			sum++
		}
	}

	return sum
}

func (s *Solution) Run1(inputPath string) int {
	reports := s.readData(inputPath)

	sum := 0
	for _, r := range reports {
		if s.isSafeWithDampener(r) {
			sum++
		}
	}

	return sum
}

func (s *Solution) isSafe(report []int) bool {
	increasing := report[1] > report[0]

	for i := 0; i < len(report)-1; i++ {
		lower, greater := report[i], report[i+1]
		if !increasing {
			lower, greater = greater, lower
		}
		if greater-lower < 1 || greater-lower > 3 {
			return false
		}
	}

	return true
}

func (s *Solution) isSafeWithDampener(report []int) bool {
	if len(report) < 3 {
		return true
	}

	a := report[1] - report[0]
	b := report[2] - report[1]
	c := report[2] - report[0]

	if len(report) == 3 {
		return a >= -3 && a <= -1 || a >= 1 && a <= 3 ||
			b >= -3 && b <= -1 || b >= 1 && b <= 3 ||
			c >= -3 && c <= -1 || c >= 1 && c <= 3
	}

	d := 0
	for _, k := range []int{a, b, c, report[3] - report[2], report[3] - report[1]} {
		if k > 0 {
			d++
		}
		if k < 0 {
			d--
		}
	}

	if d == 0 {
		return false
	}

	increasing := d > 0

	i := 0
	badSkipped := false

	for !badSkipped && i < len(report)-2 || badSkipped && i < len(report)-1 {
		a := report[i+1] - report[i]

		if !increasing {
			a = -a
		}

		if a < 1 || a > 3 {
			if badSkipped {
				return false
			}

			var b, c int

			if i > 0 {
				b = report[i+1] - report[i-1]
			}
			c = report[i+2] - report[i]
			if !increasing {
				b, c = -b, -c
			}
			if c >= 1 && c <= 3 {
				i += 2 // skip next
				badSkipped = true
				continue
			}
			if i == 0 || b >= 1 && b <= 3 {
				i++ // skip current
				badSkipped = true
				continue
			}

			return false
		}

		i++
	}

	return true
}

func (s *Solution) readData(inputPath string) [][]int {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(data), "\n")
	reports := make([][]int, len(rows))

	for i, row := range rows {
		levels := strings.Split(row, " ")
		reports[i] = make([]int, len(levels))
		for j, lvl := range levels {
			reports[i][j], err = strconv.Atoi(lvl)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return reports
}
