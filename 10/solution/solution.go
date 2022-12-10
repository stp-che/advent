package solution

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type observer interface {
	Observe(int, int)
}

type Solution struct {
	observer observer
}

func New(o observer) *Solution {
	return &Solution{o}
}

func (s *Solution) Run(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s.handleInput(file)
}

func (s *Solution) handleInput(file *os.File) {
	d := newDevice(s.observer)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d.handleCommand(scanner.Text())
	}
}

const (
	cmdNoopId = "noop"
	cmdAddxId = "addx"
)

type state struct {
	x int
}

type device struct {
	cycle    int
	state    *state
	observer observer
}

func newDevice(o observer) *device {
	return &device{
		state:    &state{x: 1},
		observer: o,
	}
}

func (d *device) handleCommand(cmd string) {
	c := parseCommand(cmd)
	c.exec(d.state, func() {
		d.cycle++
		d.observer.Observe(d.cycle, d.state.x)
	})
}

func parseCommand(cmd string) command {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case cmdNoopId:
		return cmdNoop{}
	case cmdAddxId:
		val, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return cmdAddx{int(val)}
	default:
		return nil
	}
}

type command interface {
	exec(*state, func())
}

type cmdNoop struct{}

func (c cmdNoop) exec(_ *state, eachCycle func()) {
	eachCycle()
}

type cmdAddx struct {
	val int
}

func (c cmdAddx) exec(s *state, eachCycle func()) {
	eachCycle()
	eachCycle()
	s.x += c.val
}
