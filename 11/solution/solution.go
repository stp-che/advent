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
	rounds      int
	reliefCoeff int
}

func New(reliefCoeff, rounds int) *Solution {
	return &Solution{reliefCoeff: reliefCoeff, rounds: rounds}
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
	scanner := bufio.NewScanner(file)
	monkeys := make([]*monkey, 0)
	for {
		m := parseMonkey(scanner)
		if m == nil {
			break
		}

		monkeys = append(monkeys, m)
	}

	g := newGame(s.reliefCoeff, monkeys)
	for i := 0; i < s.rounds; i++ {
		g.round()
	}
	for i, m := range g.monkeys {
		fmt.Printf("%d: %d\n", i, m.itemsInspected)
	}
	return g.monkeyBusinessLvl()
}

func parseMonkey(scanner *bufio.Scanner) *monkey {
	if !scanner.Scan() {
		return nil
	}

	m := monkey{}

	scanner.Scan()
	m.items = parseItems(scanner.Text())

	scanner.Scan()
	m.operation = parseOperation(scanner.Text())

	scanner.Scan()
	m.divider = parseDivider(scanner.Text())

	scanner.Scan()
	s1 := scanner.Text()
	scanner.Scan()
	s2 := scanner.Text()
	m.action = parseAction(s1, s2)

	scanner.Scan()

	return &m
}

func parseItems(str string) []int {
	buf := str[len("  Starting items: "):]

	itemsStr := strings.Split(buf, ", ")
	items := make([]int, len(itemsStr))
	for i, s := range itemsStr {
		items[i] = parseInt(s)
	}

	return items
}

func parseOperation(str string) operation {
	var op, id1, id2 string
	fmt.Sscanf(str, "  Operation: new = %s %s %s", &id1, &op, &id2)

	eval := func(id string, old int) int {
		if id == "old" {
			return old
		}

		return parseInt(id)
	}

	return func(old int) int {
		v1, v2 := eval(id1, old), eval(id2, old)
		switch op {
		case "*":
			return v1 * v2
		case "+":
			return v1 + v2
		default:
			log.Fatalf("Unknown operation: %s", op)
			return 0
		}
	}
}

func parseDivider(str string) int {
	var divider int
	fmt.Sscanf(str, "  Test: divisible by %d", &divider)

	return divider
}

func parseAction(ss ...string) action {
	mapping := make(map[string]int)
	for _, s := range ss {
		var key string
		var val int
		fmt.Sscanf(s, "    If %s throw to monkey %d", &key, &val)
		mapping[key] = val
	}

	return func(b bool) int {
		if b {
			return mapping["true:"]
		}

		return mapping["false:"]
	}
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(i)
}

type game struct {
	monkeys      []*monkey
	reliefCoeff  int
	dividersMult int
}

func newGame(reliefCoeff int, monkeys []*monkey) *game {
	return &game{reliefCoeff: reliefCoeff, monkeys: monkeys}
}

func (g *game) round() {
	for _, m := range g.monkeys {
		// fmt.Printf("Monkey %d\n", i)
		m.makeTurn(g.reliefCoeff, func(item, throwTo int) {
			if g.reliefCoeff == 1 && item > g.divider() {
				item = item % g.divider()
			}
			g.monkeys[throwTo].receiveItem(item)
		})
	}
}

func (g *game) divider() int {
	if g.dividersMult == 0 {
		g.dividersMult = 1
		for _, m := range g.monkeys {
			g.dividersMult *= m.divider
		}
	}

	return g.dividersMult
}

func (g *game) monkeyBusinessLvl() int {
	var f, s int
	for _, m := range g.monkeys {
		if m.itemsInspected > f {
			f, s = m.itemsInspected, f
			continue
		}
		if m.itemsInspected > s {
			s = m.itemsInspected
		}
	}

	return f * s
}

type operation func(int) int
type test func(int) bool
type action func(bool) int

type monkey struct {
	items          []int
	operation      operation
	divider        int
	action         action
	itemsInspected int
}

func (m *monkey) makeTurn(reliefCoeff int, f func(int, int)) {
	for _, i := range m.items {
		worryLevel := m.operation(i) / reliefCoeff
		// fmt.Printf("  %d => %d => %d\n", i, m.operation(i), worryLevel)
		testRes := m.test(worryLevel)
		throwTo := m.action(testRes)
		// fmt.Printf("  test(%d): %v, throwTo: %d\n", worryLevel, testRes, throwTo)
		f(worryLevel, throwTo)
	}
	m.itemsInspected += len(m.items)
	m.items = m.items[:0]
}

func (m *monkey) test(i int) bool {
	// fmt.Printf("    %d %% %d = %d\n", i, m.divider, i%m.divider)
	return i%m.divider == 0
}

func (m *monkey) receiveItem(i int) {
	m.items = append(m.items, i)
}
