package solution

type vertex struct {
	id, x, y int
}

type edge struct {
	v1, v2 vertex
	pair   int
}

func (e *edge) enterDir() [2]int {
	if e.v1.x == e.v2.x {
		if e.v1.y < e.v2.y {
			return dDown
		}

		return dUp
	}

	if e.v1.x < e.v2.x {
		return dLeft
	}

	return dRight
}

type cube struct {
	layout [5][5]byte
	// vxRanks [5][5]int
	// vxIds [6][6]int
	// clockwise list of layout border edges
	// (each edge is defined by starting vertex coordinates on layout)
	border [14]*edge //[14][2]int
	// bordersConn [14]int
	borderCalculated bool
}

func newCube() *cube {
	return &cube{}
}

func (c *cube) addSide(x, y int) {
	c.layout[x][y] = 1
}

func (c *cube) getBorder() [14]*edge {
	if !c.borderCalculated {
		vxIds := c.layoutVerticies()
		px, py := 0, 0
		for vxIds[px][py] == -1 {
			py++
		}

		i := 0
		dir := 0

		checkDir := func() bool {
			d := directions[dir]
			x, y := px+d[0], py+d[1]

			if x < 0 || x >= 6 || y < 0 || y >= 6 || vxIds[x][y] == -1 {
				return false
			}

			// fmt.Printf("x: %d, y: %d, i: %d\n", x, y, i)
			c.border[i] = &edge{
				v1: vertex{vxIds[px][py], px, py},
				v2: vertex{vxIds[x][y], x, y},
			}

			i++
			px, py = x, y
			dir--
			if dir < 0 {
				dir = 3
			}

			return true
		}

		for i < 14 {
			for {
				if checkDir() {
					break
				}

				dir++
				if dir == 4 {
					dir = 0
				}
			}
		}

		for i := 0; i < 14; i++ {
			for j := i + 1; j < 14; j++ {
				edge1, edge2 := c.border[i], c.border[j]
				b1, e1 := edge1.v1.id, edge1.v2.id
				b2, e2 := edge2.v1.id, edge2.v2.id
				if b1 == b2 && e1 == e2 || b1 == e2 && e1 == b2 {
					edge1.pair = j
					edge2.pair = i
				}
			}
		}

		c.borderCalculated = true
	}

	return c.border
}

func (c *cube) layoutVerticies() [6][6]int {
	var vxIds [6][6]int
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			vxIds[i][j] = -1
		}
	}

	visited := [5][5]bool{}

	visit := func(x, y int, vxs [8]int) {
		vxIds[x+1][y] = vxs[0]
		vxIds[x][y] = vxs[1]
		vxIds[x][y+1] = vxs[2]
		vxIds[x+1][y+1] = vxs[3]

		visited[x][y] = true
	}

	var dfs func(int, int, [8]int)
	dfs = func(x, y int, vxs [8]int) {
		if x < 0 || x > 4 || y < 0 || y > 4 {
			return
		}

		// fmt.Printf("x: %d, y: %d\n%v\n", x, y, vxs)

		visit(x, y, vxs)

		for i, p := range [4][2]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
			x, y := p[0], p[1]
			if x < 0 || x > 4 || y < 0 || y > 4 || c.layout[x][y] == 0 || visited[x][y] {
				continue
			}

			dfs(x, y, rotateCube(vxs, i))
		}
	}

	y := 0
	for c.layout[0][y] == 0 {
		y++
	}

	dfs(0, y, [8]int{0, 1, 2, 3, 4, 5, 6, 7})

	return vxIds
}

func (c *cube) goOverEdge(x1, y1, x2, y2 int) (from *edge, to *edge) {
	border := c.getBorder()
	for _, e := range border {
		if e.v1.x == x1 && e.v1.y == y1 && e.v2.x == x2 && e.v2.y == y2 {
			return e, border[e.pair]
		}
	}

	return nil, nil
}

var (
	sideFront = [4]int{0, 1, 2, 3}
	sideUp    = [4]int{1, 5, 6, 2}
	sideDown  = [4]int{4, 0, 3, 7}
	sideRight = [4]int{3, 2, 6, 7}
	sideLeft  = [4]int{4, 5, 1, 0}
)

// dir:
//  0 - up
//  1 - down
//  2 - left
//  3 - right
func rotateCube(cb [8]int, dir int) [8]int {
	switch dir {
	case 0, 1:
		cb = rotateSide(cb, sideRight, dir == 0)
		cb = rotateSide(cb, sideLeft, dir == 1)
	case 2, 3:
		cb = rotateSide(cb, sideUp, dir == 2)
		cb = rotateSide(cb, sideDown, dir == 3)
	}

	return cb
}

func rotateSide(cb [8]int, side [4]int, clockwise bool) [8]int {
	shift := 1
	if clockwise {
		shift = -1
	}

	rotatedSide := [4]int{}
	for i := 0; i < 4; i++ {
		j := i + shift
		if j < 0 {
			j = 3
		}
		if j > 3 {
			j = 0
		}
		rotatedSide[i] = cb[side[j]]
	}

	for i := 0; i < 4; i++ {
		cb[side[i]] = rotatedSide[i]
	}

	return cb
}
