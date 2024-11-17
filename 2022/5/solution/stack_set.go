package solution

type cratesMover func(n int, from, to []byte) (newFrom, newTo []byte)

type stackSet struct {
	stacks             [][]byte
	moveCratesStrategy cratesMover
}

func newStackSet(size int, mv cratesMover) *stackSet {
	s := stackSet{}
	s.stacks = make([][]byte, size)
	s.moveCratesStrategy = mv

	return &s
}

func (s *stackSet) add(stack int, crate byte) {
	s.stacks[stack] = append(s.stacks[stack], crate)
}

func (s *stackSet) move(n, from, to int) {
	newFrom, newTo := s.moveCratesStrategy(n, s.stacks[from], s.stacks[to])
	s.stacks[to] = newTo
	s.stacks[from] = newFrom
}

func (s *stackSet) topCrates() string {
	top := make([]byte, len(s.stacks))
	for i, stack := range s.stacks {
		top[i] = stack[0]
	}
	return string(top)
}
