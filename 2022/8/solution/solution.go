package solution

import (
	"bytes"
	"log"
	"os"
)

type Solution struct {
	handleInput func([][]byte) int
}

func New(handleInput func([][]byte) int) *Solution {
	return &Solution{handleInput}
}

func (s *Solution) Run(inputPath string) int {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	grid := bytes.Split(data, []byte("\n"))

	return s.handleInput(grid)
}
