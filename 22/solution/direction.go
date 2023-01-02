package solution

type direction [2]int

var (
	dRight     = direction{0, 1}
	dDown      = direction{1, 0}
	dLeft      = direction{0, -1}
	dUp        = direction{-1, 0}
	directions = []direction{dRight, dDown, dLeft, dUp}
)
