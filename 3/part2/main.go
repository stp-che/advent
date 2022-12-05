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
	elvenGroupSize     = 3
)

func priority(b byte) int {
	if b <= upperCaseZ {
		return upperCasePriorityA + int(b-upperCaseA)
	}

	return lowerCasePriorityA + int(b-lowerCaseA)
}

type elvenGroup struct {
	rucksacks       [][]byte
	currentRucksack map[byte]bool
}

func (g *elvenGroup) add(b byte) {
	if g.currentRucksack == nil {
		g.currentRucksack = make(map[byte]bool)
	}
	g.currentRucksack[b] = true
}

func (g *elvenGroup) completeRucksack() {
	rucksack := make([]byte, 0, len(g.currentRucksack))
	for b := range g.currentRucksack {
		rucksack = append(rucksack, b)
	}
	g.rucksacks = append(g.rucksacks, rucksack)
	g.currentRucksack = nil
}

func (g *elvenGroup) badge() byte {
	items := make(map[byte]int)
	for _, r := range g.rucksacks {
		for _, b := range r {
			items[b]++
			if items[b] == elvenGroupSize {
				return b
			}
		}
	}
	return 0
}

func (g *elvenGroup) priority() int {
	return priority(g.badge())
}

func handleInput(input []byte) {
	priorityTotal := 0
	group, offset := nextGroup(input, 0)
	for group != nil {
		priorityTotal += group.priority()
		group, offset = nextGroup(input, offset)
	}
	fmt.Println(priorityTotal)
}

func nextGroup(input []byte, offset int) (*elvenGroup, int) {
	if offset >= len(input) {
		return nil, 0
	}

	group := &elvenGroup{}
	for i := 0; i < elvenGroupSize; i++ {
		for ; offset < len(input) && input[offset] != '\n'; offset++ {
			group.add(input[offset])
		}
		group.completeRucksack()
		offset++
	}

	return group, offset
}
