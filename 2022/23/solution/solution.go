package solution

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) RunNRounds(inputPath string, n int) int {
	sim := s.readData(inputPath)

	// for _, e := range sim.elves {
	// 	fmt.Println(e)
	// }

	// sim.print()
	for i := 0; i < n; i++ {
		sim.round()
		// sim.print()
	}

	return sim.coveredGround()
}

func (s *Solution) RunTillStable(inputPath string) int {
	sim := s.readData(inputPath)

	rounds := 1
	for sim.round() {
		rounds++
	}

	// fmt.Println(sim.coveredGround())

	return rounds
}

func (s *Solution) readData(inputPath string) *simulation {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	sim := newSimulation()

	dataRows := bytes.Split(data, []byte{byte('\n')})
	w := len(dataRows[0])
	for i := 0; i < len(dataRows); i++ {
		for j := 0; j < w; j++ {
			if dataRows[i][j] == bOccupied {
				sim.addElf(i, j)
			}
		}
	}

	return sim
}

const bOccupied = byte('#')

type elf struct {
	x, y         int
	nextX, nextY int
}

func (e *elf) set(x, y int) {
	e.x, e.y = x, y
}

func (e *elf) makeDecision(d *disposition, lookAroundOrder [4][3][2]int) bool {
	total := 0
	firstFree := -1
	for i, cells := range lookAroundOrder {
		sum := 0
		for _, cell := range cells {
			if d.isOccupied(e.x+cell[0], e.y+cell[1]) {
				sum++
			}
		}
		if sum == 0 && firstFree == -1 {
			firstFree = i
		}
		total += sum
	}

	if total == 0 || firstFree == -1 {
		e.nextX, e.nextY = e.x, e.y
		return false
	}

	delta := lookAroundOrder[firstFree][1]

	e.nextX, e.nextY = e.x+delta[0], e.y+delta[1]

	return true
}

func (e *elf) move() {
	e.x, e.y = e.nextX, e.nextY
}

type disposition struct {
	ps map[string]int
}

func newDisposition() *disposition {
	return &disposition{
		ps: make(map[string]int),
	}
}

func (d *disposition) set(x, y, id int) {
	d.ps[key(x, y)] = id
}

func (d *disposition) isOccupied(x, y int) bool {
	_, present := d.ps[key(x, y)]

	return present
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

type simulation struct {
	elves           []*elf
	d               *disposition
	lookAroundOrder [4][3][2]int
}

func newSimulation() *simulation {
	return &simulation{
		elves: make([]*elf, 0, 4),
		d:     newDisposition(),
		lookAroundOrder: [4][3][2]int{
			{{-1, -1}, {-1, 0}, {-1, 1}},
			{{1, -1}, {1, 0}, {1, 1}},
			{{-1, -1}, {0, -1}, {1, -1}},
			{{-1, 1}, {0, 1}, {1, 1}},
		},
	}
}

func (s *simulation) addElf(x, y int) {
	s.elves = append(s.elves, &elf{x: x, y: y})
	s.d.set(x, y, len(s.elves)-1)
}

func (s *simulation) round() bool {
	anyMove := false
	nextD := newDisposition()
	choices := map[string]int{}
	for _, e := range s.elves {
		anyMove = e.makeDecision(s.d, s.lookAroundOrder) || anyMove
		choices[key(e.nextX, e.nextY)]++
	}
	for i, e := range s.elves {
		// fmt.Printf("%v", e)
		if choices[key(e.nextX, e.nextY)] == 1 {
			// fmt.Println(" - move")
			e.move()
		}
		nextD.set(e.x, e.y, i)
	}
	s.d = nextD
	s.lookAroundOrder = [4][3][2]int{
		s.lookAroundOrder[1],
		s.lookAroundOrder[2],
		s.lookAroundOrder[3],
		s.lookAroundOrder[0],
	}

	return anyMove
}

func (s *simulation) coveredGround() int {
	minX, minY, maxX, maxY := s.coveredRegion()

	return (maxX-minX+1)*(maxY-minY+1) - len(s.elves)
}

func (s *simulation) coveredRegion() (int, int, int, int) {
	minX := s.elves[0].x
	maxX := minX
	minY := s.elves[0].y
	maxY := minY

	elvesCount := len(s.elves)

	for i := 1; i < elvesCount; i++ {
		e := s.elves[i]
		if minX > e.x {
			minX = e.x
		}
		if maxX < e.x {
			maxX = e.x
		}
		if minY > e.y {
			minY = e.y
		}
		if maxY < e.y {
			maxY = e.y
		}
	}

	return minX, minY, maxX, maxY
}

func (s *simulation) print() {
	minX, minY, maxX, maxY := s.coveredRegion()

	fmt.Println()
	fmt.Printf("(%d, %d) - (%d, %d)\n----------------------\n", minX, minY, maxX, maxY)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			c := "."
			if s.d.isOccupied(i, j) {
				c = "#"
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
	fmt.Println()
}
