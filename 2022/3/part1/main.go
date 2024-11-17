package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	handleInput(input)
}

const (
	lowerCaseA         = byte('a')
	upperCaseA         = byte('A')
	upperCaseZ         = byte('Z')
	lowerCasePriorityA = 1
	upperCasePriorityA = 27
)

func priority(b byte) int {
	if b <= upperCaseZ {
		return upperCasePriorityA + int(b-upperCaseA)
	}

	return lowerCasePriorityA + int(b-lowerCaseA)
}

type rucksackStat struct {
	mins map[byte]int
	maxs map[byte]int
	size int
}

func newRucksackStat() *rucksackStat {
	rs := rucksackStat{}
	rs.mins = make(map[byte]int)
	rs.maxs = make(map[byte]int)

	return &rs
}

func (rs *rucksackStat) add(b byte) {
	if _, ok := rs.mins[b]; !ok {
		rs.mins[b] = rs.size
	}
	rs.maxs[b] = rs.size
	rs.size++
}

func (rs *rucksackStat) priority() int {
	res := 0
	half := rs.size / 2
	for b, min := range rs.mins {
		if min < half && rs.maxs[b] >= half {
			res += priority(b)
		}
	}

	return res
}

func handleInput(input []byte) {
	priorityTotal := 0
	stat, offset := nextRucksack(input, 0)
	for stat != nil {
		priorityTotal += stat.priority()
		stat, offset = nextRucksack(input, offset)
	}
	fmt.Println(priorityTotal)
}

func nextRucksack(input []byte, offset int) (*rucksackStat, int) {
	if offset >= len(input) {
		return nil, 0
	}

	stat := newRucksackStat()
	for ; offset < len(input) && input[offset] != '\n'; offset++ {
		stat.add(input[offset])
	}
	return stat, offset + 1
}
