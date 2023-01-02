package solution

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateSide(t *testing.T) {
	cb := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	cb1 := rotateSide(cb, sideFront, true)
	require.Equal(t, [8]int{3, 0, 1, 2, 4, 5, 6, 7}, cb1)

	cb1 = rotateSide(cb, sideFront, false)
	require.Equal(t, [8]int{1, 2, 3, 0, 4, 5, 6, 7}, cb1)
}

func TestLayoutVerticies(t *testing.T) {
	testCases := []struct {
		layout      [6][2]int
		expectedVxs [6][6]int
	}{
		{
			//  #
			// ####
			//   #
			layout: [6][2]int{
				{0, 1}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {2, 2},
			},
			expectedVxs: [6][6]int{
				{-1, 1, 2, -1, -1, -1},
				{1, 0, 3, 2, 1, -1},
				{5, 4, 7, 6, 5, -1},
				{-1, -1, 4, 5, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
			},
		},
		{
			// ###
			//   ###
			layout: [6][2]int{
				{0, 0}, {0, 1}, {0, 2}, {1, 2}, {1, 3}, {1, 4},
			},
			expectedVxs: [6][6]int{
				{1, 2, 6, 5, -1, -1},
				{0, 3, 7, 4, 5, 6},
				{-1, -1, 3, 0, 1, 2},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			cb := newCube()
			for _, s := range tc.layout {
				cb.addSide(s[0], s[1])
			}

			vxIds := cb.layoutVerticies()

			require.Equal(t, tc.expectedVxs, vxIds)
		})
	}
}

func TestGetBorder(t *testing.T) {
	testCases := []struct {
		layout         [6][2]int
		expectedBorder [14]*edge
	}{
		{
			//  #
			// ####
			//   #
			layout: [6][2]int{
				{0, 1}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {2, 2},
			},
			expectedBorder: [14]*edge{
				&edge{
					v1:   vertex{1, 0, 1},
					v2:   vertex{2, 0, 2},
					pair: 3,
				},
				&edge{
					v1:   vertex{2, 0, 2},
					v2:   vertex{3, 1, 2},
					pair: 2,
				},
				&edge{
					v1:   vertex{3, 1, 2},
					v2:   vertex{2, 1, 3},
					pair: 1,
				},
				&edge{
					v1:   vertex{2, 1, 3},
					v2:   vertex{1, 1, 4},
					pair: 0,
				},
				&edge{
					v1:   vertex{1, 1, 4},
					v2:   vertex{5, 2, 4},
					pair: 11,
				},
				&edge{
					v1:   vertex{5, 2, 4},
					v2:   vertex{6, 2, 3},
					pair: 6,
				},
				&edge{
					v1:   vertex{6, 2, 3},
					v2:   vertex{5, 3, 3},
					pair: 5,
				},
				&edge{
					v1:   vertex{5, 3, 3},
					v2:   vertex{4, 3, 2},
					pair: 10,
				},
				&edge{
					v1:   vertex{4, 3, 2},
					v2:   vertex{7, 2, 2},
					pair: 9,
				},
				&edge{
					v1:   vertex{7, 2, 2},
					v2:   vertex{4, 2, 1},
					pair: 8,
				},
				&edge{
					v1:   vertex{4, 2, 1},
					v2:   vertex{5, 2, 0},
					pair: 7,
				},
				&edge{
					v1:   vertex{5, 2, 0},
					v2:   vertex{1, 1, 0},
					pair: 4,
				},
				&edge{
					v1:   vertex{1, 1, 0},
					v2:   vertex{0, 1, 1},
					pair: 13,
				},
				&edge{
					v1:   vertex{0, 1, 1},
					v2:   vertex{1, 0, 1},
					pair: 12,
				},
			},
		},
	}

	borderFormatted := func(border [14]*edge) string {
		items := make([]string, 14)
		for i, e := range border {
			items[i] = fmt.Sprintf("%v", e)
		}

		return "[" + strings.Join(items, ", ") + "]"
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			cb := newCube()
			for _, s := range tc.layout {
				cb.addSide(s[0], s[1])
			}

			border := cb.getBorder()
			require.True(t,
				reflect.DeepEqual(border, tc.expectedBorder),
				fmt.Sprintf("expected: %#v\nactual: %#v", borderFormatted(tc.expectedBorder), borderFormatted(border)),
			)
		})
	}
}
