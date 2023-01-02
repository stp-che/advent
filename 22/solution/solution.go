package solution

import (
	"bytes"
	"log"
	"os"
)

type Solution struct {
	cubeSize int
}

func New(cubeSize int) *Solution {
	return &Solution{cubeSize}
}

func (s *Solution) Run(inputPath string) int {
	brd, instructions := s.readData(inputPath)
	if s.cubeSize > 0 {
		brd = brd.asCube(s.cubeSize)
	}
	actr := &actor{}
	n := 0

	for brd.isEmpty(actr.curRow, actr.curCol) {
		actr.curCol++
	}

	for _, b := range instructions {
		if b != bL && b != bR {
			n = 10*n + int(b-b0)
			continue
		}

		actr.setPosAndDir(brd.nextTile(actr.curRow, actr.curCol, n, actr.getDir()))
		n = 0
		actr.turn(b)
	}

	if n > 0 {
		actr.setPosAndDir(brd.nextTile(actr.curRow, actr.curCol, n, actr.getDir()))
	}

	return 1000*(actr.curRow+1) + 4*(actr.curCol+1) + actr.curDir
}

func (s *Solution) readData(inputPath string) (*board, []byte) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	dataRows := bytes.Split(data, []byte{byte('\n')})
	n := len(dataRows)

	return newBoard(dataRows[:n-2]), dataRows[n-1]
}

// const (
// 	tileEmpty = byte(' ')
// 	tileOpen  = byte('.')
// 	tileWall  = byte('#')

// 	b0 = byte('0')
// 	bL = byte('L')
// 	bR = byte('R')
// )

// type board [][]byte

// func (b board) nextTile(row, col, dist int, dir [2]int) (int, int) {
// 	var nextRow, nextCol int
// 	for i := 0; i < dist; i++ {
// 		nextRow, nextCol = row+dir[0], col+dir[1]
// 		if b.isEmpty(nextRow, nextCol) {
// 			nextRow, nextCol = b.goToOppositeSide(row, col, dir)
// 		}

// 		if b[nextRow][nextCol] == tileWall {
// 			break
// 		}

// 		row, col = nextRow, nextCol
// 	}

// 	return row, col
// }

// func (b board) isEmpty(row, col int) bool {
// 	return row < 0 || row >= b.height() || col < 0 || col >= b.width(row) || b[row][col] == tileEmpty
// }

// func (b board) goToOppositeSide(row, col int, dir [2]int) (int, int) {
// 	switch dir {
// 	case dRight:
// 		col = 0
// 	case dDown:
// 		row = 0
// 	case dLeft:
// 		col = b.width(row) - 1
// 	case dUp:
// 		row = b.height() - 1
// 	}

// 	for b.isEmpty(row, col) {
// 		row += dir[0]
// 		col += dir[1]
// 	}

// 	return row, col
// }

// func (b board) height() int {
// 	return len(b)
// }

// func (b board) width(row int) int {
// 	return len(b[row])
// }

// var (
// 	dRight     = [2]int{0, 1}
// 	dDown      = [2]int{1, 0}
// 	dLeft      = [2]int{0, -1}
// 	dUp        = [2]int{-1, 0}
// 	directions = [][2]int{dRight, dDown, dLeft, dUp}
// )

// type actor struct {
// 	curRow, curCol, curDir int
// }

// func (a *actor) setPos(row, col int) {
// 	a.curRow, a.curCol = row, col
// }

// func (a *actor) turn(dirCode byte) {
// 	d := 1
// 	if dirCode == bL {
// 		d = -1
// 	}

// 	a.curDir += d

// 	if a.curDir == -1 {
// 		a.curDir = len(directions) - 1
// 	}

// 	if a.curDir == len(directions) {
// 		a.curDir = 0
// 	}
// }

// func (a *actor) getDir() [2]int {
// 	return directions[a.curDir]
// }
