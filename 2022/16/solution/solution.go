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

func (s *Solution) Run(inputPath string, t, actors int) int {
	cave := s.readData(inputPath)

	return s.calculate(cave, t, actors)
}

func (s *Solution) calculate(cave *cave, t, actors int) int {
	max := 0
	var bests [][]int

	eachUniqSeparation(cave.usefulValves(), actors, func(parts []map[int]bool) {
		opns := make([][]int, actors)
		res := 0
		for i, part := range parts {
			x, o := cave.pressureReleased(cave.start, t, part)
			res += x
			opns[i] = o
		}

		if res > max {
			max = res
			bests = opns
		}
	})

	for _, b := range bests {
		fmt.Println(b)
	}

	return max
}

func (s *Solution) readData(inputPath string) *cave {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	cave := newCaveParsing(string(data)).result()

	return cave
}

type caveParsing struct {
	mapping map[string]int
	cave    *cave
	data    []string
}

func newCaveParsing(data string) *caveParsing {
	return &caveParsing{
		mapping: make(map[string]int),
		data:    strings.Split(data, "\n"),
	}
}

func (p *caveParsing) result() *cave {
	n := len(p.data)

	p.cave = &cave{
		dist:  make([][]int, n),
		rates: make([]int, n),
	}

	p.parseValves()
	p.calcDist()

	p.cave.start = p.mapping["AA"]

	fmt.Println(p.mapping)

	return p.cave
}

func (p *caveParsing) parseValves() {
	n := len(p.data)

	for _, s := range p.data {
		id, rate, connections := parseValve(s)
		i := p.gatValveN(id)

		p.cave.rates[i] = rate

		p.cave.dist[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			p.cave.dist[i][j] = -1
		}
		for _, c := range connections {
			p.cave.dist[i][p.gatValveN(c)] = 1
		}
	}
}

func (p *caveParsing) calcDist() {
	n := len(p.data)
	d := 1

	for {
		new := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if p.cave.dist[i][j] == d {
					for k := i + 1; k < n; k++ {
						if p.cave.dist[j][k] == 1 && p.cave.dist[i][k] < 0 {
							new++
							p.cave.dist[i][k] = d + 1
							p.cave.dist[k][i] = d + 1
						}
					}
				}
			}
		}

		d++
		if new == 0 {
			break
		}
	}
}

func (p *caveParsing) gatValveN(id string) int {
	if _, ok := p.mapping[id]; !ok {
		p.mapping[id] = len(p.mapping)
	}

	return p.mapping[id]
}

const (
	valvePart1Pattern    = "Valve %s has flow rate=%d"
	valvePart2PrefixOne  = " tunnel leads to valve "
	valvePart2PrefixMany = " tunnels lead to valves "
)

func parseValve(s string) (string, int, []string) {
	var id string
	var rate int

	parts := strings.Split(s, ";")
	fmt.Sscanf(parts[0], valvePart1Pattern, &id, &rate)

	return id, rate, parseValvePart2(parts[1])
}

func parseValvePart2(s string) []string {
	l := len(valvePart2PrefixMany)
	if strings.HasPrefix(s, valvePart2PrefixOne) {
		l = len(valvePart2PrefixOne)
	}

	return strings.Split(s[l:], ", ")
}

type cave struct {
	dist  [][]int
	rates []int
	start int
}

func (c *cave) usefulValves() []int {
	res := []int{}

	for i, r := range c.rates {
		if r > 0 {
			res = append(res, i)
		}
	}

	return res
}

func (c *cave) pressureReleased(start, minutes int, valves map[int]bool) (int, []int) {
	max := 0
	var best []int
	for v, opened := range valves {
		if opened {
			continue
		}

		minRest := minutes - c.dist[start][v] - 1
		if minRest <= 0 {
			continue
		}

		valves[v] = true
		res, opn := c.pressureReleased(v, minRest, valves)
		valves[v] = false

		res += minRest * c.rates[v]

		if res > max {
			max = res
			best = append([]int{v}, opn...)
		}
	}

	return max, best
}

func eachUniqSeparation(list []int, parts int, f func([]map[int]bool)) {
	var rec func(int, int, int, []map[int]bool)

	rec = func(pts, fix, n int, collected []map[int]bool) {
		if pts == 1 {
			part := make(map[int]bool)
			j := 1
			for _, v := range list {
				if fix&j == 0 {
					part[v] = false
				}
				j = j << 1
			}

			f(append(collected, part))

			return
		}

		nn := 1 << n
		for i := 1; i < nn; i += 2 {
			part := make(map[int]bool)

			newFix := fix
			newN := n

			k := 1
			x := 1
			for _, v := range list {
				if fix&x == 0 {
					if i&k != 0 {
						part[v] = false
						newFix = newFix | x
						newN--
					}

					k = k << 1
				}

				x = x << 1
			}

			rec(pts-1, newFix, newN, append(collected, part))
		}
	}

	rec(parts, 0, len(list), []map[int]bool{})
}
