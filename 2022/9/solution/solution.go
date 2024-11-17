package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	ropeLength int
}

func New(ropeLength int) *Solution {
	return &Solution{ropeLength}
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
	g := newRope(s.ropeLength)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		n, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		g.moveHead(parts[0], int(n))
	}

	return g.tailVisitedCount()
}

type cell struct {
	x, y int
}

func (c cell) String() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}

func (c cell) diff(other cell) (int, int) {
	return c.x - other.x, c.y - other.y
}

type rope struct {
	knots     []cell
	tailTrace map[string]bool
}

func newRope(lenght int) *rope {
	r := rope{
		knots:     make([]cell, lenght),
		tailTrace: make(map[string]bool),
	}
	r.tailTrace[r.knots[lenght-1].String()] = true

	return &r
}

func (r *rope) tailVisitedCount() int {
	return len(r.tailTrace)
}

func (r *rope) moveHead(dir string, n int) {
	dx, dy := delta(dir)
	for i := 0; i < n; i++ {
		r.knots[0].x += dx
		r.knots[0].y += dy
		r.pullTail()
	}
}

func (r *rope) pullTail() {
	for i := 1; i < len(r.knots); i++ {
		dx, dy := r.knots[i-1].diff(r.knots[i])
		if -2 < dx && dx < 2 && -2 < dy && dy < 2 {
			return
		}

		if dx > 0 {
			r.knots[i].x++
		}
		if dx < 0 {
			r.knots[i].x--
		}
		if dy > 0 {
			r.knots[i].y++
		}
		if dy < 0 {
			r.knots[i].y--
		}
	}
	r.tailTrace[r.knots[len(r.knots)-1].String()] = true
}

func delta(dir string) (int, int) {
	switch dir {
	case "U":
		return 0, 1
	case "D":
		return 0, -1
	case "R":
		return 1, 0
	case "L":
		return -1, 0
	default:
		log.Fatalf("Wring direction: %s", dir)
		return 0, 0
	}
}
