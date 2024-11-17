package solution

import (
	"bytes"
	"log"
	"os"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) ShortestPathFromStart(inputPath string) int {
	return s.readInput(inputPath).minPathLenFromStart()
}

func (s *Solution) ShortestPathFromAnyA(inputPath string) int {
	return s.readInput(inputPath).minPathLenFromAnyA()
}

func (s *Solution) readInput(inputPath string) *area {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return newArea(bytes.Split(data, []byte("\n")))
}

type cell struct {
	x, y int
}

type area struct {
	heightmap [][]byte
	start     cell
	target    cell
	pathLens  [][]int
}

func newArea(data [][]byte) *area {
	a := area{heightmap: data}
	rows, cols := len(data), len(data[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch data[i][j] {
			case byte('S'):
				a.start = cell{i, j}
				a.heightmap[i][j] = byte('a')
			case byte('E'):
				a.target = cell{i, j}
				a.heightmap[i][j] = byte('z')
			}
		}
	}

	return &a
}

func (a *area) minPathLenFromStart() int {
	a.calculatePathLens()

	return a.pathLens[a.start.x][a.start.y]
}

func (a *area) minPathLenFromAnyA() int {
	a.calculatePathLens()

	rows, cols := len(a.pathLens), len(a.pathLens[0])
	min := a.pathLens[a.start.x][a.start.y]
	elevation := byte('a')
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if a.heightmap[i][j] == elevation && a.pathLens[i][j] >= 0 && min > a.pathLens[i][j] {
				min = a.pathLens[i][j]
			}
		}
	}
	return min
}

func (a *area) calculatePathLens() {
	cells := []cell{a.target}
	d := []int{-1, 0, 1, 0, -1}
	rows, cols := len(a.heightmap), len(a.heightmap[0])
	a.pathLens = make([][]int, rows)
	for i := 0; i < rows; i++ {
		a.pathLens[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			a.pathLens[i][j] = -1
		}
	}
	a.pathLens[a.target.x][a.target.y] = 0

	for len(cells) > 0 {
		nextCells := []cell{}
		for _, c := range cells {
			for i := 0; i < 4; i++ {
				x, y := c.x+d[i], c.y+d[i+1]
				if x < 0 || x >= rows || y < 0 || y >= cols || a.pathLens[x][y] >= 0 {
					continue
				}

				h1, h2 := a.heightmap[c.x][c.y], a.heightmap[x][y]

				if h1 > h2+1 {
					continue
				}

				nextCells = append(nextCells, cell{x, y})
				a.pathLens[x][y] = a.pathLens[c.x][c.y] + 1
			}
		}
		cells = nextCells
	}

	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < cols; j++ {
	// 		color := "\033[37m"
	// 		if a.pathLens[i][j] >= 0 {
	// 			color = "\033[33m"
	// 		}
	// 		fmt.Print(color)
	// 		fmt.Print(string([]byte{a.heightmap[i][j]}))
	// 	}
	// 	fmt.Println()
	// }
}
