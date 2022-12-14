package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type itemObserver interface {
}

type Solution struct {
	withFloorAt int
}

func New(withFloorAt int) *Solution {
	return &Solution{withFloorAt}
}

func (s *Solution) Run(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return s.handleInput(file)
}

func (s *Solution) handleInput(file *os.File) int {
	paths := make([]path, 0, 4)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		paths = append(paths, parsePath(scanner.Text()))
	}

	rs := newRockSlice(point{500, 0}, paths, s.withFloorAt)

	res := rs.fillWithSand()

	for _, row := range rs.grid {
		for _, t := range row {
			var c string
			switch t {
			case cAir:
				c = "."
			case cRock:
				c = "#"
			case cSand:
				c = "o"
			}
			fmt.Print(c)
		}
		fmt.Println()
	}

	return res
}

func parsePath(s string) path {
	pointsStr := strings.Split(s, " -> ")
	p := make(path, len(pointsStr))
	for i, ps := range pointsStr {
		p[i] = parsePoint(ps)
	}

	return p
}

func parsePoint(s string) point {
	xy := strings.Split(s, ",")

	return point{
		parseInt(xy[0]),
		parseInt(xy[1]),
	}
}

func parseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(n)
}

type point struct {
	x, y int
}

type path []point

func (p path) bounds() (point, point) {
	min := p[0]
	max := p[0]

	for i := 1; i < len(p); i++ {
		if min.x > p[i].x {
			min.x = p[i].x
		}
		if min.y > p[i].y {
			min.y = p[i].y
		}
		if max.x < p[i].x {
			max.x = p[i].x
		}
		if max.y < p[i].y {
			max.y = p[i].y
		}
	}

	return min, max
}

type cellType byte

const (
	cAir cellType = iota
	cRock
	cSand
)

type rockSlice struct {
	grid       [][]cellType
	sandSource point
	minX       int
}

func newRockSlice(sandSource point, paths []path, withFloorAt int) *rockSlice {
	min, max := getGridBounds(paths)
	if withFloorAt > 0 {
		min, max = adjustGridBoundsWithFloor(min, max, sandSource, withFloorAt)
		paths = append(paths, path{
			point{min.x, max.y},
			point{max.x, max.y},
		})
	}
	grid := make([][]cellType, max.y+1)
	width := max.x - min.x + 1
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]cellType, width)
	}

	s := &rockSlice{
		sandSource: sandSource,
		grid:       grid,
		minX:       min.x,
	}

	for _, p := range paths {
		s.addPath(p)
	}

	return s
}

func (s *rockSlice) get(x, y int) cellType {
	return s.grid[y][x-s.minX]
}

func (s *rockSlice) set(x, y int, t cellType) {
	s.grid[y][x-s.minX] = t
}

func (s *rockSlice) addPath(p path) {
	for i := 0; i < len(p)-1; i++ {
		p1, p2 := p[i], p[i+1]
		if p1.x == p2.x {
			y1, y2 := p1.y, p2.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for ; y1 <= y2; y1++ {
				s.set(p1.x, y1, cRock)
			}

			continue
		}

		x1, x2 := p1.x, p2.x
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for ; x1 <= x2; x1++ {
			s.set(x1, p1.y, cRock)
		}
	}
}

func (s *rockSlice) fillWithSand() int {
	var x, y int
	var rest bool
	count := 0
	for {
		rest, x, y = s.nextSandPart()

		if !rest {
			break
		}

		// fmt.Printf("%v %d %d\n", rest, x, y)
		s.set(x, y, cSand)
		count++

		if y == s.sandSource.y {
			break
		}
	}

	return count
}

func (s *rockSlice) nextSandPart() (bool, int, int) {
	x, y := s.sandSource.x-s.minX, s.sandSource.y
	rows, cols := len(s.grid), len(s.grid[0])
	stop := false
	for !stop && y < rows-1 && x > 0 && x < cols-1 {
		if s.grid[y+1][x] == cAir {
			y++
			continue
		}

		if s.grid[y+1][x-1] == cAir {
			y++
			x--
			continue
		}

		if s.grid[y+1][x+1] == cAir {
			y++
			x++
			continue
		}

		stop = true
	}

	return y < rows-1 && x > 0 && x < cols-1, x + s.minX, y
}

func getGridBounds(paths []path) (point, point) {
	min, max := paths[0].bounds()
	for _, p := range paths {
		pMin, pMax := p.bounds()
		if min.x > pMin.x {
			min.x = pMin.x
		}
		if min.y > pMin.y {
			min.y = pMin.y
		}
		if max.x < pMax.x {
			max.x = pMax.x
		}
		if max.y < pMax.y {
			max.y = pMax.y
		}
	}

	return min, max
}

func adjustGridBoundsWithFloor(min, max, sandSource point, withFloorAt int) (point, point) {
	max.y = max.y + withFloorAt
	h := max.y - sandSource.y

	minX := sandSource.x - h
	if min.x > minX {
		min.x = minX
	}

	maxX := sandSource.x + h
	if max.x < maxX {
		max.x = maxX
	}

	return min, max
}
