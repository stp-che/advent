package solution

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	operands []int
	value    int
}

func (e Equation) isValid(useConcat bool) bool {
	var check func(int, []int) bool
	check = func(acc int, rest []int) bool {
		if acc > e.value {
			return false
		}
		if len(rest) == 0 {
			return acc == e.value
		}

		a1 := acc + rest[0]
		a2 := acc * rest[0]
		var a3 int
		if useConcat {
			m := 10
			for m <= rest[0] {
				m *= 10
			}
			a3 = acc*m + rest[0]
		}
		rest = rest[1:]
		return check(a1, rest) || check(a2, rest) || useConcat && check(a3, rest)
	}

	return check(e.operands[0], e.operands[1:])
}

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	equations := s.readData(inputPath)
	sum := 0
	for _, e := range equations {
		// fmt.Println(e.value)
		// fmt.Println(e.operands)
		if e.isValid(false) {
			sum += e.value
		}
	}
	return sum
}

func (s *Solution) Run1(inputPath string) int {
	equations := s.readData(inputPath)
	sum := 0
	for _, e := range equations {
		// fmt.Println(e.value)
		// fmt.Println(e.operands)
		if e.isValid(true) {
			sum += e.value
		}
	}
	return sum
}

func (s *Solution) readData(inputPath string) []*Equation {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	res := make([]*Equation, len(lines))
	for i, line := range lines {
		res[i] = parseEquation(line)
	}

	return res
}

func parseEquation(str string) *Equation {
	e := Equation{}
	parts := strings.Split(str, ": ")
	e.value, _ = strconv.Atoi(parts[0])
	parts = strings.Split(parts[1], " ")
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		e.operands = append(e.operands, n)
	}
	return &e
}
