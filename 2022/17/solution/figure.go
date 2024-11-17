package solution

type figure struct {
	top int
	// widthBound byte // deprecated
	// lines      []byte // deprecated
	shifts   [][]byte
	curShift int
	_height  int
}

func newFigure(width, top, left int, form *form) *figure {
	shifts := make([][]byte, width-form.width()+1)
	for i := 0; i < len(shifts); i++ {
		shifts[i] = make([]byte, form.height())
		for j, formLine := range form.lines {
			shifts[i][j] = byte(formLine << (len(shifts) - i - 1))
		}
	}
	// shift := width - form.width() - left
	// lines := make([]byte, form.height())
	// for i, formLine := range form.lines {
	// 	lines[i] = byte(formLine << shift)
	// }

	return &figure{
		top:      top,
		curShift: left,
		shifts:   shifts,
		_height:  form.height(),

		// widthBound: byte(1 << (width - 1)),
		// lines: lines,
	}
}

func (f *figure) height() int {
	return f._height
}

func (f *figure) moveDown() {
	f.top--
}

func (f *figure) moveLeft() bool {
	if f.curShift > 0 {
		f.curShift--
		return true
	}

	return false
	// for _, l := range f.lines {
	// 	if l >= f.widthBound {
	// 		return
	// 	}
	// }

	// for i, l := range f.lines {
	// 	f.lines[i] = l << 1
	// }
}

func (f *figure) moveRight() bool {
	if f.curShift < len(f.shifts)-1 {
		f.curShift++
		return true
	}

	return false
	// for _, l := range f.lines {
	// 	if l&1 == 1 {
	// 		return
	// 	}
	// }

	// for i, l := range f.lines {
	// 	f.lines[i] = l >> 1
	// }
}

func (f *figure) lines() []byte {
	return f.shifts[f.curShift]
}

type cycleFiguresGenerator struct {
	figuresList []*figure
	current     int
}

func newCycleFiguressGenerator(width int, forms [][]byte) *cycleFiguresGenerator {
	figuresList := make([]*figure, len(forms))
	for i, s := range forms {
		f := newForm(s)
		figuresList[i] = newFigure(width, f.height(), 0, f)
	}

	return &cycleFiguresGenerator{
		figuresList: figuresList,
	}
}

func (g *cycleFiguresGenerator) nextFigure() *figure {
	// defer func() {
	// 	g.current = (g.current + 1) % len(g.figuresList)
	// }()
	i := g.current
	g.current++
	if g.current == len(g.figuresList) {
		g.current = 0
	}

	return g.figuresList[i]
}
