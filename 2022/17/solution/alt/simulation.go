package alt

import (
	"fmt"
)

const (
	wLeft  = byte('<')
	wRight = byte('>')
)

type formsGenerator interface {
	nextForm() Form
}

const (
	newFigGap  = 3
	newFigLeft = 2
)

type Simulation struct {
	width       int
	newFigGap   int
	newFigLeft  int
	winds       []byte
	curWind     int
	fGen        formsGenerator
	stack       [][]byte
	_verbose    bool
	stackOffset int
	maxes       []int
}

func NewSimulation(width, newFigGap, newFigLeft int, winds []byte, fGen formsGenerator) *Simulation {
	return &Simulation{
		width:      width,
		newFigGap:  newFigGap,
		newFigLeft: newFigLeft,
		winds:      winds,
		fGen:       fGen,
		maxes:      make([]int, width),
	}
}

func (s *Simulation) StackHeight() int {
	return s.stackOffset + len(s.stack)
}

func (s *Simulation) Verbose() *Simulation {
	s._verbose = true
	return s
}

func (s *Simulation) DropNextFigure() {
	f := s.fGen.nextForm()
	top := len(s.stack) + newFigGap + f.height() - 1
	fig := newFigure(top, newFigLeft, f)

	s.printStack(fig)

	s.dropFigure(fig)
	s.addToStack(fig)
	// s.cutStack()

	s.printStack(nil)
}

func (s *Simulation) dropFigure(fig *figure) {
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

func (s *Simulation) blowWind(fig *figure) {
	defer func() {
		s.curWind = (s.curWind + 1) % len(s.winds)
	}()

	shift := func() {
		if fig.left > 0 {
			fig.left--
		}
	}
	rollback := func() {
		if fig.left+fig.width() < s.width {
			fig.left++
		}
	}

	if s._verbose {
		fmt.Println(string([]byte{s.winds[s.curWind]}))
	}

	if s.winds[s.curWind] == wRight {
		shift, rollback = rollback, shift
	}

	shift()

	if intersects(fig.top, fig.left, fig.form, s.stack) {
		rollback()
	}
}

func (s *Simulation) canMoveDown(fig *figure) bool {
	return fig.top >= fig.height() && !intersects(fig.top-1, fig.left, fig.form, s.stack)
}

func (s *Simulation) addToStack(fig *figure) {
	if fig.top >= len(s.stack) {
		for i := len(s.stack); i <= fig.top; i++ {
			s.stack = append(s.stack, make([]byte, s.width))
		}
	}

	w := fig.width()

	for i, line := range fig.form {
		n := fig.top - i
		for j := 0; j < w; j++ {
			if line[j] == 1 {
				s.stack[n][fig.left+j] = 1
			}
		}
	}
}

// func (s *Simulation) cutStack() {
// 	i := s.availableDepth()
// 	s.stackOffset += i
// 	s.stack = s.stack[i:]
// }

// func (s *Simulation) availableDepth() int {
// 	widthMask := byte((1 << s.width) - 1)
// 	above := widthMask

// 	if s._verbose {
// 		fmt.Println()
// 	}

// 	for i := len(s.stack) - 1; i >= 0; i-- {
// 		x := ^s.stack[i] & widthMask
// 		cur := above & x

// 		if cur == 0 {
// 			return i
// 		}

// 		for j := 0; j < s.width; j++ {
// 			cur = cur | ((cur<<1)|(cur>>1))&x
// 		}

// 		if s._verbose {
// 			for j := s.width - 1; j >= 0; j-- {
// 				c := "O"
// 				if cur&(1<<j) == 0 {
// 					c = "."
// 				}
// 				fmt.Print(c)
// 			}
// 			fmt.Println()
// 		}

// 		above = cur
// 	}

// 	return 0
// }

func (s *Simulation) printStack(fig *figure) {
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
		line := make([]byte, s.width)
		fLine := make([]byte, s.width)

		if i < len(s.stack) {
			line = s.stack[i]
		}

		if fig != nil && i <= fig.top && i > fig.top-fig.height() {
			fLine = fig.form[fig.top-i]
		}

		for j := 0; j < s.width; j++ {
			switch {
			case fig != nil && j >= fig.left && j-fig.left < fig.width() && fLine[j-fig.left] == 1:
				c = "@"
			case line[j] == 1:
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

func intersects(top, left int, form Form, stack [][]byte) bool {
	bottom := top - form.height() + 1

	if bottom >= len(stack) {
		return false
	}

	w := form.width()

	for i := bottom; i < len(stack) && i <= top; i++ {
		for j := 0; j < w; j++ {
			if stack[i][left+j] == 1 && form[top-i][j] == 1 {
				return true
			}
		}
	}

	return false
}
