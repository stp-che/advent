package solution

type form struct {
	lines []byte
	w     int
}

func newForm(lines []byte) *form {
	return &form{
		lines: lines,
	}
}

func (f *form) height() int {
	return len(f.lines)
}

func (f *form) width() int {
	if f.w == 0 {
		for _, l := range f.lines {
			for (1 << f.w) <= l {
				f.w++
			}
		}
	}

	return f.w
}

var (
	formHorLine  = []byte{0b1111}
	formPlus     = []byte{0b010, 0b111, 0b010}
	formAngle    = []byte{0b001, 0b001, 0b111}
	formVertLine = []byte{1, 1, 1, 1}
	formBox      = []byte{0b11, 0b11}
)

type cycleFormsGenerator struct {
	formsList []*form
	current   int
}

func newCycleFormsGenerator(forms [][]byte) *cycleFormsGenerator {
	formsList := make([]*form, len(forms))
	for i, s := range forms {
		formsList[i] = newForm(s)
	}

	return &cycleFormsGenerator{
		formsList: formsList,
	}
}

func (g *cycleFormsGenerator) nextForm() *form {
	defer func() {
		g.current = (g.current + 1) % len(g.formsList)
	}()

	return g.formsList[g.current]
}
