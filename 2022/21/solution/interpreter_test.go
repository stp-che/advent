package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpressionRevertWithParam(t *testing.T) {
	pName := "foo"

	cases := []struct {
		expr        string            //expression
		ctx         map[string]string //map[string]expression
		x           int
		expectedRes int
	}{
		{
			expr:        pName, //&identifier{pName},
			x:           10,
			expectedRes: 10,
		},
		{
			expr: "foo + aaa",
			ctx: map[string]string{
				"aaa": "7",
			},
			x:           10,
			expectedRes: 3,
		},
		{
			expr: "aaa + foo",
			ctx: map[string]string{
				"aaa": "7",
			},
			x:           10,
			expectedRes: 3,
		},
		{
			expr: "foo - aaa",
			ctx: map[string]string{
				"aaa": "7",
			},
			x:           3,
			expectedRes: 10,
		},
		{
			expr: "aaa - foo",
			ctx: map[string]string{
				"aaa": "7",
			},
			x:           3,
			expectedRes: 4,
		},
		{
			expr: "foo * aaa",
			ctx: map[string]string{
				"aaa": "7",
			},
			x:           35,
			expectedRes: 5,
		},
		{
			expr: "aaa * foo",
			ctx: map[string]string{
				"aaa": "7",
			},
			x:           35,
			expectedRes: 5,
		},
		{
			expr: "foo / aaa",
			ctx: map[string]string{
				"aaa": "7",
			},
			x:           5,
			expectedRes: 35,
		},
		{
			expr: "aaa / foo",
			ctx: map[string]string{
				"aaa": "42",
			},
			x:           6,
			expectedRes: 7,
		},
		{
			expr: "a / b",
			ctx: map[string]string{
				"a": "foo + c",
				"c": "5",
				"b": "2",
			},
			x:           10,
			expectedRes: 15,
		},
		{
			expr: "cczh / lfqf",
			ctx: map[string]string{
				"cczh": "sllz + lgvd",
				"lfqf": "4",
				"ljgn": "2",
				"sllz": "4",
				"lgvd": "ljgn * foo",
			},
			x:           150,
			expectedRes: 298,
		},
		{
			expr: "cczh / lfqf",
			ctx: map[string]string{
				"cczh": "sllz + lgvd",
				"ptdq": "foo - dvpt",
				"dvpt": "3",
				"lfqf": "4",
				"ljgn": "2",
				"sllz": "4",
				"lgvd": "ljgn * ptdq",
			},
			x:           150,
			expectedRes: 301,
		},
		{
			expr: "cczh / lfqf",
			ctx: map[string]string{
				"root": "pppw + sjmn",
				"dbpl": "5",
				"cczh": "sllz + lgvd",
				"zczc": "2",
				"ptdq": "foo - dvpt",
				"dvpt": "3",
				"lfqf": "4",
				"foo":  "5",
				"ljgn": "2",
				"sjmn": "drzm * dbpl",
				"sllz": "4",
				"pppw": "cczh / lfqf",
				"lgvd": "ljgn * ptdq",
				"drzm": "hmdt - zczc",
				"hmdt": "32",
			},
			x:           150,
			expectedRes: 301,
		},
	}

	parser := newParser()

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			expr := parser.parseExpression(c.expr)
			ctx := map[string]expression{}
			for k, e := range c.ctx {
				ctx[k] = parser.parseExpression(e)
			}

			f, err := expr.revertWithParam(ctx, pName)
			require.NoError(t, err)
			res := f(c.x)
			require.Equal(t, c.expectedRes, res)
		})
	}

}
