package solution

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	cubes := s.readData(inputPath)

	surface := len(cubes) * 6

	for i := 0; i < len(cubes); i++ {
		for j := i + 1; j < len(cubes); j++ {
			if i == j {
				continue
			}

			if connected(cubes[i], cubes[j]) {
				surface = surface - 2
			}
		}
	}

	return surface

}

func (s *Solution) Run1(inputPath string) int {
	cubes := s.readData(inputPath)

	return calcSurface(collect(cubes))
}

func (s *Solution) readData(inputPath string) [][3]int {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(data), "\n")
	res := make([][3]int, len(rows))

	for i, row := range rows {
		res[i] = [3]int{}
		fmt.Sscanf(row, "%d,%d,%d", &res[i][0], &res[i][1], &res[i][2])
	}

	return res
}

func collect(cubes [][3]int) [][][]bool {
	maxes := [3]int{}
	mins := [3]int{-1, -1, -1}

	for _, c := range cubes {
		for j := 0; j < 3; j++ {
			if maxes[j] < c[j] {
				maxes[j] = c[j]
			}
			if mins[j] == -1 || mins[j] > c[j] {
				mins[j] = c[j]
			}
		}
	}

	deltas := [3]int{mins[0] - 1, mins[1] - 1, mins[2] - 1}

	res := make([][][]bool, maxes[0]+2)
	for i := 0; i < maxes[0]-deltas[0]+2; i++ {
		res[i] = make([][]bool, maxes[1]-deltas[1]+2)
		for j := 0; j < maxes[1]-deltas[1]+2; j++ {
			res[i][j] = make([]bool, maxes[2]-deltas[2]+2)
		}
	}

	for _, c := range cubes {
		res[c[0]-deltas[0]][c[1]-deltas[1]][c[2]-deltas[2]] = true
	}

	return res
}

func calcSurface(collected [][][]bool) int {
	lenX, lenY, lenZ := len(collected), len(collected[0]), len(collected[0][0])
	visited := make([][][]bool, lenX)
	for i := 0; i < lenX; i++ {
		visited[i] = make([][]bool, lenY)
		for j := 0; j < lenY; j++ {
			visited[i][j] = make([]bool, lenZ)
		}
	}

	next := [][3]int{{lenX - 1, lenY - 1, lenZ - 1}}
	visited[lenX-1][lenY-1][lenZ-1] = true

	surface := 0

	for len(next) > 0 {
		newNext := [][3]int{}
		for _, c := range next {
			eachConnected(c[0], c[1], c[2], func(x, y, z int) {
				if !(0 <= x && x < lenX && 0 <= y && y < lenY && 0 <= z && z < lenZ) {
					return
				}

				if visited[x][y][z] {
					return
				}

				if collected[x][y][z] {
					surface++
					return
				}

				visited[x][y][z] = true
				newNext = append(newNext, [3]int{x, y, z})
			})

			next = newNext
		}
	}

	return surface
}

var connectedRelative = [][3]int{
	{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1},
}

func eachConnected(x, y, z int, fn func(int, int, int)) {
	for _, c := range connectedRelative {
		fn(x+c[0], y+c[1], z+c[2])
	}
}

func connected(a, b [3]int) bool {
	x := a[0] - b[0]
	y := a[1] - b[1]
	z := a[2] - b[2]

	return x*x+y*y+z*z == 1
}
