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

func (s *Solution) Run(inputPath string) int {
	sim := s.readData(inputPath)

	return sim.bfs(sim.start, sim.target)
}

func (s *Solution) Run1(inputPath string) int {
	sim := s.readData(inputPath)

	steps := sim.bfs(sim.start, sim.target)
	steps += sim.bfs(sim.target, sim.start)
	steps += sim.bfs(sim.start, sim.target)

	return steps
}

// func (s *Solution) RunTillStable(inputPath string) int {
// 	sim := s.readData(inputPath)

// 	rounds := 1
// 	for sim.round() {
// 		rounds++
// 	}

// 	// fmt.Println(sim.coveredGround())

// 	return rounds
// }

func (s *Solution) readData(inputPath string) *simulation {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	sim := newSimulation(bytes.Split(data, []byte{byte('\n')}))

	return sim
}

const (
	bEmpty    = byte('.')
	bDirUp    = byte('^')
	bDirRight = byte('>')
	bDirDown  = byte('v')
	bDirLeft  = byte('<')
	bBusy     = byte('*')
	bWall     = byte('#')
)

type blizzard struct {
	x, y   int
	dx, dy int
}

func newBlizzard(x, y int, dir byte) *blizzard {
	var dx, dy int
	switch dir {
	case bDirUp:
		dx = -1
	case bDirRight:
		dy = 1
	case bDirDown:
		dx = 1
	case bDirLeft:
		dy = -1
	}

	return &blizzard{x, y, dx, dy}
}

func (b *blizzard) move(w, h int) {
	b.x += b.dx
	b.y += b.dy
	if b.x == 0 {
		b.x = h - 2
	}
	if b.x == h-1 {
		b.x = 1
	}
	if b.y == 0 {
		b.y = w - 2
	}
	if b.y == w-1 {
		b.y = 1
	}
}

type simulation struct {
	h, w      int
	cave      [][]byte
	start     [2]int
	target    [2]int
	blizzards []*blizzard
}

func newSimulation(dataRows [][]byte) *simulation {
	var start, target [2]int
	blizzards := make([]*blizzard, 0, 4)

	h, w := len(dataRows), len(dataRows[0])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			c := dataRows[i][j]
			if i == 0 && c == bEmpty {
				start = [2]int{i, j}
			}
			if i == h-1 && c == bEmpty {
				target = [2]int{i, j}
			}

			if c != bEmpty && c != bWall {
				blizzards = append(blizzards, newBlizzard(i, j, c))
			}
		}
	}

	return &simulation{
		h:         h,
		w:         w,
		cave:      dataRows,
		start:     start,
		target:    target,
		blizzards: blizzards,
	}
}

func (s *simulation) addBlizzard(b *blizzard) {
	s.blizzards = append(s.blizzards, b)
}

func (s *simulation) bfs(start, target [2]int) int {
	cells := [][2]int{start}
	step := 0
	stepVisited := make([][]bool, s.h)
	for i := 0; i < s.h; i++ {
		stepVisited[i] = make([]bool, s.w)
	}
	ds := [][2]int{{0, 0}, {-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for {
		nextCells := [][2]int{}
		s.moveBlizzards()
		step++

		for _, c := range cells {
			for _, d := range ds {
				x, y := c[0]+d[0], c[1]+d[1]
				if x == target[0] && y == target[1] {
					return step
				}

				if x < 0 || x >= s.h || y < 0 || y >= s.w || stepVisited[x][y] {
					continue
				}

				if s.cave[x][y] == bEmpty {
					stepVisited[x][y] = true
					nextCells = append(nextCells, [2]int{x, y})
				}
			}
		}

		for _, c := range cells {
			stepVisited[c[0]][c[1]] = false
		}
		for _, c := range nextCells {
			stepVisited[c[0]][c[1]] = false
		}

		cells = nextCells
	}
}

func (s *simulation) moveBlizzards() {
	for i := 1; i < s.h-1; i++ {
		for j := 1; j < s.w-1; j++ {
			s.cave[i][j] = bEmpty
		}
	}

	for _, b := range s.blizzards {
		b.move(s.w, s.h)
		s.cave[b.x][b.y] = bBusy
	}
}
