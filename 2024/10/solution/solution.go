package solution

import (
	"bytes"
	"log"
	"os"
)

var (
	bottom = byte('0')
	top    = byte('9')
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	theMap := s.readData(inputPath)
	sum := 0

	for i, row := range theMap {
		for j, c := range row {
			if c == bottom {
				sum += s.topsRating(theMap, i, j)
			}
		}
	}

	return sum
}

func (s *Solution) Run1(inputPath string) int {
	theMap := s.readData(inputPath)
	sum := 0

	for i, row := range theMap {
		for j, c := range row {
			if c == bottom {
				sum += s.trailsRating(theMap, i, j)
			}
		}
	}

	return sum
}

func (s *Solution) topsRating(aMap [][]byte, i, j int) int {
	height := len(aMap)
	width := len(aMap[0])
	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	shifts := [5]int{-1, 0, 1, 0, -1}

	toVisit := [][2]int{{i, j}}
	offset := 0
	topsVisited := 0
	for offset < len(toVisit) {
		for nextOffset := len(toVisit); offset < nextOffset; offset++ {
			i := toVisit[offset][0]
			j := toVisit[offset][1]
			if visited[i][j] {
				continue
			}

			visited[i][j] = true

			if aMap[i][j] == top {
				topsVisited++
				continue
			}

			for k := 0; k < 4; k++ {
				ii := i + shifts[k]
				jj := j + shifts[k+1]
				if ii < 0 || ii >= height || jj < 0 || jj >= width || visited[ii][jj] {
					continue
				}
				if aMap[ii][jj] == aMap[i][j]+1 {
					toVisit = append(toVisit, [2]int{ii, jj})
				}
			}
		}

	}

	return topsVisited
}

func (s *Solution) trailsRating(aMap [][]byte, i, j int) int {
	height := len(aMap)
	width := len(aMap[0])

	shifts := [5]int{-1, 0, 1, 0, -1}

	toVisit := [][2]int{{i, j}}
	offset := 0
	topsVisited := 0
	for offset < len(toVisit) {
		for nextOffset := len(toVisit); offset < nextOffset; offset++ {
			i := toVisit[offset][0]
			j := toVisit[offset][1]

			if aMap[i][j] == top {
				topsVisited++
				continue
			}

			for k := 0; k < 4; k++ {
				ii := i + shifts[k]
				jj := j + shifts[k+1]
				if ii < 0 || ii >= height || jj < 0 || jj >= width {
					continue
				}
				if aMap[ii][jj] == aMap[i][j]+1 {
					toVisit = append(toVisit, [2]int{ii, jj})
				}
			}
		}

	}

	return topsVisited
}

func (s *Solution) readData(inputPath string) [][]byte {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{byte('\n')})
}
