package solution

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	movePattern = "move %d from %d to %d"
)

func New(mv cratesMover) *Solution {
	return &Solution{mv}
}

type Solution struct {
	moveCratesStrategy cratesMover
}

func (s *Solution) Run(inputPath string) string {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	return s.handleInput(input)
}

func (s *Solution) handleInput(input []byte) string {
	rows := strings.Split(string(input), "\n")
	stacks, offset := s.getInitialStacks(rows)
	for ; offset < len(rows); offset++ {
		n, from, to := parseMove(rows[offset])
		stacks.move(n, from-1, to-1)
	}
	// fmt.Println(offset)
	// fmt.Println(stacks)
	return stacks.topCrates()
}

func (s *Solution) getInitialStacks(rows []string) (*stackSet, int) {
	stacksCount := (len(rows[0]) + 1) / 4
	// fmt.Printf("stacksCount: %d\n", stacksCount)
	stacks := newStackSet(stacksCount, s.moveCratesStrategy)
	i := 0
	for ; rows[i+1] != ""; i++ {
		// fmt.Printf("i: %d\n", i)
		for j := 0; j < stacksCount; j++ {
			crate := rows[i][j*4+1]
			// fmt.Printf("i: %d, j: %d, j*4+1: %d, crate: %s\n", i, j, j*4+1, string(crate))
			if crate == ' ' {
				continue
			}

			stacks.add(j, crate)
		}
	}
	return stacks, i + 2
}

func parseMove(row string) (int, int, int) {
	var n, from, to int
	_, err := fmt.Sscanf(row, movePattern, &n, &from, &to)
	if err != nil {
		log.Fatal(err)
	}

	return n, from, to
}
