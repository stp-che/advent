package solution

import (
	"fmt"
)

const (
	wLeft  = byte('<')
	wRight = byte('>')
)

// type formsGenerator interface {
// 	nextForm() *form
// }

type figuresGenerator interface {
	nextFigure() *figure
}

const (
	newFigGap  = 3
	newFigLeft = 2
)

type simulation struct {
	width       int
	newFigGap   int
	newFigLeft  int
	winds       []byte
	curWind     int
	fGen        figuresGenerator
	stack       []byte
	_verbose    bool
	stackOffset int
	maxes       []int
}

func newSimulation(width, newFigGap, newFigLeft int, winds []byte, fGen figuresGenerator) *simulation {
	return &simulation{
		width:      width,
		newFigGap:  newFigGap,
		newFigLeft: newFigLeft,
		winds:      winds,
		fGen:       fGen,
		maxes:      make([]int, width),
	}
}

func (s *simulation) stackHeight() int {
	return s.stackOffset + len(s.stack)
}

func (s *simulation) verbose() *simulation {
	s._verbose = true
	return s
}

func (s *simulation) dropNextFigure() {
	// f := s.fGen.nextFfigure()
	// top := len(s.stack) + newFigGap + f.height() - 1
	// fig := newFigure(s.width, top, newFigLeft, f)
	fig := s.fGen.nextFigure()
	fig.top = len(s.stack) + newFigGap + fig.height() - 1
	fig.curShift = newFigLeft

	// s.printStack(fig)

	s.dropFigure(fig)
	s.addToStack(fig)
	// s.cutStack()

	// s.printStack(nil)
}

func (s *simulation) dropFigure(fig *figure) {
	for {
		s.blowWind(fig)

		s.printStack(fig)

		if !s.canMoveDown(fig) {
			break
		}

		fig.moveDown()

		s.printStack(fig)
	}
}

func (s *simulation) blowWind(fig *figure) {
	wind := s.winds[s.curWind]
	s.curWind++
	if s.curWind == len(s.winds) {
		s.curWind = 0
	}

	shift := fig.moveLeft
	rollback := fig.moveRight

	if s._verbose {
		fmt.Println(string([]byte{s.winds[s.curWind]}))
	}

	if wind == wRight {
		shift, rollback = rollback, shift
	}

	if !shift() {
		return
	}

	if intersects(fig.top, fig.shifts[fig.curShift], s.stack) {
		rollback()
	}
}

func (s *simulation) canMoveDown(fig *figure) bool {
	return fig.top >= fig.height() && !intersects(fig.top-1, fig.shifts[fig.curShift], s.stack)
}

func (s *simulation) addToStack(fig *figure) {
	if fig.top >= len(s.stack) {
		for i := len(s.stack); i <= fig.top; i++ {
			s.stack = append(s.stack, 0)
		}
	}

	for i, line := range fig.shifts[fig.curShift] {
		n := fig.top - i
		s.stack[n] = s.stack[n] | line
	}
}

func (s *simulation) cutStack() {
	i := s.availableDepth()
	s.stackOffset += i
	s.stack = s.stack[i:]
}

func (s *simulation) availableDepth() int {
	widthMask := byte((1 << s.width) - 1)
	above := widthMask

	if s._verbose {
		fmt.Println()
	}

	for i := len(s.stack) - 1; i >= 0; i-- {
		x := ^s.stack[i] & widthMask
		cur := above & x

		if cur == 0 {
			return i
		}

		for j := 0; j < s.width; j++ {
			cur = cur | ((cur<<1)|(cur>>1))&x
		}

		if s._verbose {
			for j := s.width - 1; j >= 0; j-- {
				c := "O"
				if cur&(1<<j) == 0 {
					c = "."
				}
				fmt.Print(c)
			}
			fmt.Println()
		}

		above = cur
	}

	return 0
}

func (s *simulation) printStack(fig *figure) {
	if !s._verbose {
		return
	}

	var c string

	fmt.Println()
	start := len(s.stack) - 1

	if fig != nil && fig.top > start {
		start = fig.top
	}

	for i := start; i >= 0; i-- {
		line := byte(0)
		fLine := byte(0)

		if i < len(s.stack) {
			line = s.stack[i]
		}

		if fig != nil && i <= fig.top && i > fig.top-fig.height() {
			fLine = fig.lines()[fig.top-i]
		}

		for j := s.width - 1; j >= 0; j-- {
			x := byte(1 << j)
			switch {
			case fLine&x > 0:
				c = "@"
			case line&x > 0:
				c = "#"
			default:
				c = "."
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func intersects(fTop int, fLines, stack []byte) bool {
	bottom := fTop - len(fLines) + 1

	if bottom >= len(stack) {
		return false
	}

	for i := bottom; i < len(stack) && i <= fTop; i++ {
		if stack[i]&fLines[fTop-i] > 0 {
			return true
		}
	}

	return false
}
