package solution

import (
	"log"
	"os"
	"slices"
	"strconv"
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
	followers, updates := s.readData(inputPath)
	// fmt.Println(followers)
	sum := 0
	for _, u := range updates {
		// fmt.Println(u)
		if s.isValid(u, followers) {
			m, _ := strconv.Atoi(u[len(u)/2])
			// fmt.Printf("    valid, middle: %d\n", m)
			sum += m
		}
	}
	return sum
}

func (s *Solution) Run1(inputPath string) int {
	followers, updates := s.readData(inputPath)
	// fmt.Println(followers)
	sum := 0
	for _, u := range updates {
		// fmt.Println(u)
		if !s.isValid(u, followers) {
			s.correct(u, followers)
			m, _ := strconv.Atoi(u[len(u)/2])
			// fmt.Printf("    valid, middle: %d\n", m)
			sum += m
		}
	}
	return sum
}

func (s *Solution) isValid(update []string, followers map[string]map[string]struct{}) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if _, ok := followers[update[j]][update[i]]; ok {
				return false
			}
		}
	}

	return true
}

func (s *Solution) correct(update []string, followers map[string]map[string]struct{}) {
	slices.SortFunc(update, func(a, b string) int {
		if _, ok := followers[a][b]; ok {
			return -1
		}
		if _, ok := followers[b][a]; ok {
			return 1
		}
		return 0
	})
}

func (s *Solution) readData(inputPath string) (map[string]map[string]struct{}, [][]string) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	followers := map[string]map[string]struct{}{}
	updates := [][]string{}

	lines := strings.Split(string(data), "\n")
	i := 0
	for lines[i] != "" {
		rule := strings.Split(lines[i], "|")
		if followers[rule[0]] == nil {
			followers[rule[0]] = map[string]struct{}{}
		}
		followers[rule[0]][rule[1]] = struct{}{}
		i++
	}

	i++
	for i < len(lines) {
		updates = append(updates, strings.Split(lines[i], ","))
		i++
	}

	return followers, updates
}
