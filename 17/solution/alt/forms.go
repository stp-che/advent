package alt

type Form [][]byte

func (f Form) height() int {
	return len(f)
}

func (f Form) width() int {
	return len(f[0])
}

var (
	FormHorLine = Form{{1, 1, 1, 1}}
	FormPlus    = Form{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}
	FormAngle = Form{
		{0, 0, 1},
		{0, 0, 1},
		{1, 1, 1},
	}
	FormVertLine = Form{{1}, {1}, {1}, {1}}
	FormBox      = Form{
		{1, 1},
		{1, 1},
	}
)

type CycleFormsGenerator struct {
	formsList []Form
	current   int
}

func NewCycleFormsGenerator(forms []Form) *CycleFormsGenerator {
	return &CycleFormsGenerator{
		formsList: forms,
	}
}

func (g *CycleFormsGenerator) nextForm() Form {
	defer func() {
		g.current = (g.current + 1) % len(g.formsList)
	}()

	return g.formsList[g.current]
}
