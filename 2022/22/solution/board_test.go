package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoOverEdge(t *testing.T) {
	b := newBoard(
		[][]byte{
			[]byte("     .....          "),
			[]byte("     .....          "),
			[]byte("     .....          "),
			[]byte("     .....          "),
			[]byte("     .....          "),
			[]byte("...................."),
			[]byte("...................."),
			[]byte("...................."),
			[]byte("...................."),
			[]byte("...................."),
			[]byte("          .....     "),
			[]byte("          .....     "),
			[]byte("          .....     "),
			[]byte("          .....     "),
			[]byte("          .....     "),
		},
	).asCube(5)

	type pos struct {
		x, y int
		dir  [2]int
	}

	testCases := []struct {
		from, to pos
	}{
		{
			from: pos{0, 6, dUp},
			to:   pos{5, 18, dDown},
		},
		{
			from: pos{4, 5, dLeft},
			to:   pos{5, 4, dDown},
		},
		{
			from: pos{9, 0, dLeft},
			to:   pos{9, 19, dLeft},
		},
		{
			from: pos{9, 0, dDown},
			to:   pos{14, 14, dUp},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			x, y, dir := b.goOverEdge(tc.from.x, tc.from.y, tc.from.dir)

			require.Equal(t, tc.to, pos{x, y, dir})
		})
	}
}
