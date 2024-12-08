package solution

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSituationFallsInCycle(t *testing.T) {
	s := newSituationFromString(
		strings.Join(
			[]string{
				".#..",
				"...#",
				"#...",
				"..#.",
			},
			"\n",
		),
	)

	require.True(t, s.fallsInCycle(0, 2, down))
	require.False(t, s.fallsInCycle(3, 3, up))

	s = newSituationFromString(
		strings.Join(
			[]string{
				".#......",
				".......#",
				"...#....",
				".....#..",
				"#.......",
				"....#...",
				"..#.....",
				"......#.",
			},
			"\n",
		),
	)

	require.False(t, s.fallsInCycle(7, 0, up))
	require.True(t, s.fallsInCycle(1, 0, right))
}
