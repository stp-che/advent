package solution

import (
	"log"
	"os"
	"strings"
)

type Map struct {
	width, height int
	antennas      map[rune][][2]int
}

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	m := s.readData(inputPath)
	antinodes := map[int]map[int]struct{}{}
	sum := 0

	check := func(h, w int) {
		if h < 0 || h >= m.height || w < 0 || w >= m.width {
			return
		}
		if _, ok := antinodes[h][w]; ok {
			return
		}
		sum++
		if antinodes[h] == nil {
			antinodes[h] = map[int]struct{}{}
		}
		antinodes[h][w] = struct{}{}
	}

	for _, antennas := range m.antennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				dh := antennas[j][0] - antennas[i][0]
				dw := antennas[j][1] - antennas[i][1]
				check(antennas[j][0]+dh, antennas[j][1]+dw)
				check(antennas[i][0]-dh, antennas[i][1]-dw)
			}
		}
	}
	return sum
}

func (s *Solution) Run1(inputPath string) int {
	m := s.readData(inputPath)
	antinodes := map[int]map[int]struct{}{}
	sum := 0

	check := func(h, w int) bool {
		if h < 0 || h >= m.height || w < 0 || w >= m.width {
			return false
		}
		if _, ok := antinodes[h][w]; ok {
			return true
		}
		sum++
		if antinodes[h] == nil {
			antinodes[h] = map[int]struct{}{}
		}
		antinodes[h][w] = struct{}{}

		return true
	}

	for _, antennas := range m.antennas {
		if len(antennas) < 2 {
			continue
		}
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				_ = check(antennas[i][0], antennas[i][1])
				_ = check(antennas[j][0], antennas[j][1])
				dh := antennas[j][0] - antennas[i][0]
				dw := antennas[j][1] - antennas[i][1]
				h := antennas[j][0] + dh
				w := antennas[j][1] + dw
				for check(h, w) {
					h += dh
					w += dw
				}
				h = antennas[i][0] - dh
				w = antennas[i][1] - dw
				for check(h, w) {
					h -= dh
					w -= dw
				}
			}
		}
	}
	return sum
}

func (s *Solution) readData(inputPath string) Map {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	m := Map{
		antennas: map[rune][][2]int{},
	}

	lines := strings.Split(string(data), "\n")
	m.height = len(lines)
	m.width = len(lines[0])
	for i, line := range lines {
		for j, c := range line {
			if c != '.' {
				m.antennas[c] = append(m.antennas[c], [2]int{i, j})
			}
		}
	}

	return m
}
