package solution

import (
	"bufio"
	"log"
	"os"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Part1(inputPath string) int {
	return s.parseData(inputPath).evalRoot()
}

func (s *Solution) Part2(inputPath, param string) int {
	return s.parseData(inputPath).calcParam(param)
}

func (s *Solution) parseData(inputPath string) *interpreter {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	parser := newParser()

	for scanner.Scan() {
		parser.parseLine(scanner.Text())
	}

	return &interpreter{parser.ctx}
}
