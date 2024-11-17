package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	cases := [][3]string{
		{"10", "1", "11"},
		{"1=", "2=", "21"},
		{"20", "10", "1=0"},
		{"10", "1=0", "1-0"},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := add([]byte(c[0]), []byte(c[1]))
			require.Equal(t, c[2], string(res))
		})
	}
}
