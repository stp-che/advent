package solution

import (
	"log"
	"os"
)

type state string

const (
	nullSt             state = ""
	mSt                state = "m"
	muSt               state = "mu"
	mulSt              state = "mul"
	dSt                state = "d"
	doSt               state = "do"
	donSt              state = "don"
	don_St             state = "don'"
	dontSt             state = "don't"
	doCallSt           state = "doCall"
	dontCallSt         state = "dontCall"
	waitForFirstArgSt  state = "waitForFirstArg"
	firstArgSt         state = "firstArg"
	waitForSecondArgSt state = "waitForSecondArg"
	secondArgSt        state = "secondArg"
	disabledSt         state = "disabled"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	input := s.readData(inputPath)
	sum := 0
	s.forEachMul(input, false, func(i1, i2 int) { sum += i1 * i2 })

	return sum
}

func (s *Solution) Run1(inputPath string) int {
	input := s.readData(inputPath)
	sum := 0
	s.forEachMul(input, true, func(i1, i2 int) { sum += i1 * i2 })

	return sum
}

func (s *Solution) forEachMul(input string, useDo bool, handleMul func(int, int)) {
	var firstArg, secondArg int
	var st state
	for _, c := range input {
		if c == 'd' && useDo {
			st = dSt
		} else if c == 'o' && st == dSt {
			st = doSt
		} else if c == 'n' && st == doSt {
			st = donSt
		} else if c == '\'' && st == donSt {
			st = don_St
		} else if c == 't' && st == don_St {
			st = dontSt
		} else if c == 'm' && st != disabledSt {
			st = mSt
		} else if c == 'u' && st == mSt {
			st = muSt
		} else if c == 'l' && st == muSt {
			st = mulSt
		} else if c == '(' {
			switch st {
			case mulSt:
				st = waitForFirstArgSt
			case doSt:
				st = doCallSt
			case dontSt:
				st = dontCallSt
			}
		} else if c == ',' && st == firstArgSt {
			st = waitForSecondArgSt
		} else if c == ')' {
			switch st {
			case secondArgSt:
				handleMul(firstArg, secondArg)
				st = nullSt
			case doCallSt:
				st = nullSt
			case dontCallSt:
				st = disabledSt
			}
		} else if '0' <= c && c <= '9' && st != disabledSt {
			switch st {
			case waitForFirstArgSt:
				st = firstArgSt
				firstArg = int(c - '0')
			case firstArgSt:
				firstArg = firstArg*10 + int(c-'0')
			case waitForSecondArgSt:
				st = secondArgSt
				secondArg = int(c - '0')
			case secondArgSt:
				secondArg = secondArg*10 + int(c-'0')
			default:
				st = nullSt
			}
		} else if st != disabledSt {
			st = nullSt
		}
	}
}

func (s *Solution) readData(inputPath string) string {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
