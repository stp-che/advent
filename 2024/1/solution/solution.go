package solution

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	left, right := s.readData(inputPath)
	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		l, r := left[i], right[i]
		if r < l {
			l, r = r, l
		}
		sum += r - l
	}

	return sum
}

func (s *Solution) Run1(inputPath string) int {
	left, right := s.readData(inputPath)

	ds := map[int]int{}
	for _, n := range right {
		ds[n]++
	}

	sum := 0
	for _, n := range left {
		sum += n * ds[n]
	}

	return sum
}

func (s *Solution) readData(inputPath string) ([]int, []int) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(data), "\n")
	left := make([]int, len(rows))
	right := make([]int, len(rows))

	for i, row := range rows {
		fmt.Sscanf(row, "%d   %d", &left[i], &right[i])
	}

	return left, right
}
