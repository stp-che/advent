package solution

const (
	tileEmpty = byte(' ')
	tileOpen  = byte('.')
	tileWall  = byte('#')

	b0 = byte('0')
	bL = byte('L')
	bR = byte('R')
)

type board struct {
	data     [][]byte
	cb       *cube
	cubeSize int
}

func newBoard(data [][]byte) *board {
	return &board{data: data}
}

func (b *board) asCube(size int) *board {
	return &board{
		data:     b.data,
		cubeSize: size,
	}
}

func (b *board) nextTile(row, col, dist int, dir [2]int) (int, int, [2]int) {
	var nextRow, nextCol int
	nextDir := dir
	for i := 0; i < dist; i++ {
		nextRow, nextCol = row+dir[0], col+dir[1]
		if b.isEmpty(nextRow, nextCol) {
			if b.cubeSize > 0 {
				nextRow, nextCol, nextDir = b.goOverEdge(row, col, dir)
			} else {
				nextRow, nextCol = b.goToOppositeSide(row, col, dir)
			}
		}

		if b.data[nextRow][nextCol] == tileWall {
			break
		}

		row, col, dir = nextRow, nextCol, nextDir
	}

	return row, col, dir
}

func (b *board) isEmpty(row, col int) bool {
	return row < 0 || row >= b.height() || col < 0 || col >= b.width(row) || b.data[row][col] == tileEmpty
}

func (b *board) goToOppositeSide(row, col int, dir [2]int) (int, int) {
	switch dir {
	case dRight:
		col = 0
	case dDown:
		row = 0
	case dLeft:
		col = b.width(row) - 1
	case dUp:
		row = b.height() - 1
	}

	for b.isEmpty(row, col) {
		row += dir[0]
		col += dir[1]
	}

	return row, col
}

func (b *board) height() int {
	return len(b.data)
}

func (b *board) width(row int) int {
	return len(b.data[row])
}

func (b *board) getCube() *cube {
	if b.cb == nil {
		b.cb = newCube()

		for i := 0; i*b.cubeSize < b.height(); i++ {
			for j := 0; j*b.cubeSize < b.width(i*b.cubeSize); j++ {
				if !b.isEmpty(i*b.cubeSize, j*b.cubeSize) {
					b.cb.addSide(i, j)
				}
			}
		}
	}

	return b.cb
}

func (b *board) goOverEdge(row, col int, dir [2]int) (int, int, [2]int) {
	x1, y1, x2, y2 := b.getEdgeEnds(row, col, dir)
	eFrom, eTo := b.getCube().goOverEdge(x1, y1, x2, y2)
	n := b.getOffsetOnEdge(row, col, eFrom)

	newRow, newCol := b.applyOffsetOnEdge(eTo, n, eFrom.v1.id != eTo.v1.id)
	// fmt.Printf("applyOffsetOnEdge(%v, %d, %v) => (%d, %d)\n", eTo, n, eFrom.v1.id != eTo.v1.id, newRow, newCol)

	return newRow, newCol, eTo.enterDir()
}

func (b *board) getEdgeEnds(row, col int, dir [2]int) (int, int, int, int) {
	x, y := row/b.cubeSize, col/b.cubeSize
	switch dir {
	case dUp:
		return x, y, x, y + 1
	case dRight:
		return x, y + 1, x + 1, y + 1
	case dDown:
		return x + 1, y + 1, x + 1, y
	case dLeft:
		return x + 1, y, x, y
	}

	return -1, -1, -1, -1
}

func (b *board) getOffsetOnEdge(row, col int, e *edge) int {
	x, y := b.getEdgeEnd(e, e.v1.x, e.v1.y)

	if row == x {
		return (col - y) * (e.v2.y - e.v1.y)
	}

	return (row - x) * (e.v2.x - e.v1.x)
}

func (b *board) applyOffsetOnEdge(e *edge, offset int, reverse bool) (int, int) {
	v1, v2 := e.v1, e.v2
	if reverse {
		v1, v2 = v2, v1
	}

	// fmt.Printf("    v1, v2: %v, %v\n", v1, v2)

	row, col := b.getEdgeEnd(e, v1.x, v1.y)
	dx, dy := v2.x-v1.x, v2.y-v1.y

	return row + dx*offset, col + dy*offset
}

func (b *board) getEdgeEnd(e *edge, x, y int) (int, int) {
	row := x * b.cubeSize
	col := y * b.cubeSize
	d := e.enterDir()
	if x > e.v1.x || x > e.v2.x || d == dUp {
		row--
	}
	if y > e.v1.y || y > e.v2.y || d == dLeft {
		col--
	}

	return row, col
}
