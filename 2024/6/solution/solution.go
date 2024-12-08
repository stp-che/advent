package solution

import (
	"log"
	"os"
	"strings"
)

var (
	directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	up         = 0
	right      = 1
	down       = 2
	left       = 3
)

var (
	obstacle = 1
	checked  = 1 << 1
	traces   = []int{1 << 2, 1 << 3, 1 << 4, 1 << 5}
)

type Situation struct {
	board [][]int
	guard [2]int
}

func (s *Situation) fallsInCycle(i, j, dir int) bool {
	// fmt.Printf("   fallsInCycle(%d, %d, %d)\n", i, j, dir)
	board := s.copyBoard()
	height := len(board)
	width := len(board[0])
	for {
		if board[i][j]&traces[dir] > 0 {
			return true
		}
		board[i][j] |= traces[dir]

		nextI := i + directions[dir][0]
		nextJ := j + directions[dir][1]
		if nextI < 0 || nextI == height || nextJ < 0 || nextJ == width {
			return false
		}
		if board[nextI][nextJ] == obstacle {
			dir = (dir + 1) % len(directions)
			continue
		}
		i = nextI
		j = nextJ
	}
}

func (s *Situation) copyBoard() [][]int {
	res := make([][]int, len(s.board))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(s.board[i]))
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = s.board[i][j]
		}
	}
	return res
}

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	st := s.readData(inputPath)
	visited := 1
	d := up
	height := len(st.board)
	width := len(st.board[0])
	for {
		i := st.guard[0] + directions[d][0]
		j := st.guard[1] + directions[d][1]
		if i < 0 || i == height || j < 0 || j == width {
			break
		}
		if st.board[i][j] == 1 {
			d = (d + 1) % len(directions)
			continue
		}
		st.guard[0] = i
		st.guard[1] = j
		if st.board[i][j] == 2 {
			continue
		}
		st.board[i][j] = 2
		visited++
	}

	return visited
}

func (s *Solution) Run1(inputPath string) int {
	st := s.readData(inputPath)
	sum := 0
	dir := up
	i, j := st.guard[0], st.guard[1]
	height := len(st.board)
	width := len(st.board[0])
	for {
		// fmt.Printf("i: %d, j: %d, dir: %d\n", i, j, dir)
		nextI := i + directions[dir][0]
		nextJ := j + directions[dir][1]
		if nextI < 0 || nextI == height || nextJ < 0 || nextJ == width {
			break
		}
		if st.board[nextI][nextJ] == obstacle {
			dir = (dir + 1) % len(directions)
			continue
		}
		if !(nextI == st.guard[0] && nextJ == st.guard[1]) && st.board[nextI][nextJ] != checked {
			st.board[nextI][nextJ] = obstacle
			if st.fallsInCycle(i, j, (dir+1)%len(directions)) {
				sum++
			}
			st.board[nextI][nextJ] = checked
		}
		i = nextI
		j = nextJ
	}
	return sum
}

func (s *Solution) readData(inputPath string) *Situation {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return newSituationFromString(string(data))
}

func newSituationFromString(data string) *Situation {
	lines := strings.Split(data, "\n")
	st := &Situation{}
	st.board = make([][]int, len(lines))
	for i, line := range lines {
		st.board[i] = make([]int, len(line))
		for j, c := range line {
			if c == '#' {
				st.board[i][j] = obstacle
			}
			if c == '^' {
				st.board[i][j] = 2
				st.guard = [2]int{i, j}
			}
		}
	}

	return st
}
