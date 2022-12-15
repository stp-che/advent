package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnionAdd(t *testing.T) {
	type testCase struct {
		u   union
		in  interval
		res union
	}

	cases := []testCase{
		{
			u:   union{},
			in:  newInterval(3, 7),
			res: union{{3, 7}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(-3, -1),
			res: union{{-3, -1}, {1, 5}, {8, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(13, 21),
			res: union{{1, 5}, {8, 10}, {13, 21}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(-1, 12),
			res: union{{-1, 12}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(1, 10),
			res: union{{1, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(3, 6),
			res: union{{1, 6}, {8, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(7, 9),
			res: union{{1, 5}, {7, 10}},
		},
		{
			u:   union{{8, 10}},
			in:  newInterval(3, 7),
			res: union{{3, 10}},
		},
		{
			u:   union{{8, 10}},
			in:  newInterval(11, 17),
			res: union{{8, 17}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(6, 7),
			res: union{{1, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}, {20, 30}},
			in:  newInterval(14, 17),
			res: union{{1, 5}, {8, 10}, {14, 17}, {20, 30}},
		},
		{
			u:   union{{1, 5}, {8, 10}, {20, 30}, {40, 50}},
			in:  newInterval(9, 25),
			res: union{{1, 5}, {8, 30}, {40, 50}},
		},
		{
			u:   union{{1, 5}, {20, 30}, {40, 50}},
			in:  newInterval(22, 25),
			res: union{{1, 5}, {20, 30}, {40, 50}},
		},
	}

	for _, c := range cases {
		actual := c.u.add(c.in)
		require.Equal(t, c.res, actual)
	}
}

func TestUnitRemove(t *testing.T) {
	type testCase struct {
		u   union
		in  interval
		res union
	}

	cases := []testCase{
		{
			u:   union{},
			in:  newInterval(3, 7),
			res: union{},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(-3, -1),
			res: union{{1, 5}, {8, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(13, 21),
			res: union{{1, 5}, {8, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(-1, 12),
			res: union{},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(1, 10),
			res: union{},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(3, 6),
			res: union{{1, 2}, {8, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(7, 9),
			res: union{{1, 5}, {10, 10}},
		},
		{
			u:   union{{8, 10}},
			in:  newInterval(3, 8),
			res: union{{9, 10}},
		},
		{
			u:   union{{8, 10}},
			in:  newInterval(10, 17),
			res: union{{8, 9}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(6, 7),
			res: union{{1, 5}, {8, 10}},
		},
		{
			u:   union{{1, 5}, {8, 10}, {20, 30}},
			in:  newInterval(4, 22),
			res: union{{1, 3}, {23, 30}},
		},
		{
			u:   union{{1, 50}},
			in:  newInterval(9, 25),
			res: union{{1, 8}, {26, 50}},
		},
		{
			u:   union{{1, 5}, {8, 10}},
			in:  newInterval(2, 9),
			res: union{{1, 1}, {10, 10}},
		},
	}

	for _, c := range cases {
		actual := c.u.remove(c.in)
		require.Equal(t, c.res, actual)
	}
}
