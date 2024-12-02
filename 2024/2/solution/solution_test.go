package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSafeWithDampener(t *testing.T) {
	s := New()
	require.Equal(t, false, s.isSafeWithDampener([]int{1, 2, 7, 8, 9}))
	require.Equal(t, false, s.isSafeWithDampener([]int{9, 7, 6, 2, 1}))
	require.Equal(t, true, s.isSafeWithDampener([]int{1, 3, 2, 4, 5}))
	require.Equal(t, true, s.isSafeWithDampener([]int{1, 8, 2, 4, 5}))
	require.Equal(t, true, s.isSafeWithDampener([]int{8, 6, 4, 4, 1}))
	require.Equal(t, true, s.isSafeWithDampener([]int{3, 1, 2, 4, 5}))
	require.Equal(t, true, s.isSafeWithDampener([]int{1, 2, 3, 4, 15}))
	require.Equal(t, true, s.isSafeWithDampener([]int{1, 15}))
	require.Equal(t, true, s.isSafeWithDampener([]int{1, 2}))
	require.Equal(t, true, s.isSafeWithDampener([]int{1, 15, 2}))
	require.Equal(t, true, s.isSafeWithDampener([]int{15, 1, 2}))
	require.Equal(t, false, s.isSafeWithDampener([]int{15, 1, 8}))
	require.Equal(t, false, s.isSafeWithDampener([]int{1, 2, 3, 2, 1}))
}
