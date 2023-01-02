package solution

type actor struct {
	curRow, curCol, curDir int
}

func (a *actor) setPosAndDir(row, col int, dir [2]int) {
	a.curRow, a.curCol = row, col
	a.curDir = 0
	for dir != directions[a.curDir] {
		a.curDir++
	}
}

func (a *actor) turn(dirCode byte) {
	d := 1
	if dirCode == bL {
		d = -1
	}

	a.curDir += d

	if a.curDir == -1 {
		a.curDir = len(directions) - 1
	}

	if a.curDir == len(directions) {
		a.curDir = 0
	}
}

func (a *actor) getDir() direction {
	return directions[a.curDir]
}
