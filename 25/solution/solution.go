package solution

import (
	"bytes"
	"log"
	"os"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) string {
	sum := []byte{zero}
	for _, n := range s.readData(inputPath) {
		sum = add(sum, n)
	}

	return string(sum)
}

func (s *Solution) readData(inputPath string) [][]byte {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{byte('\n')})
}

const (
	baseSize = 5
	minDigit = -2
	maxDigit = 2
	zero     = byte('0')
)

var (
	digits = []byte("=-012")
	dVals  = map[byte]int{
		digits[0]: -2,
		digits[1]: -1,
		digits[2]: 0,
		digits[3]: 1,
		digits[4]: 2,
	}
)

func add(a, b []byte) []byte {
	if len(a) < len(b) {
		a, b = b, a
	}

	res := make([]byte, len(a)+1)
	rest := 0
	i := 1
	for ; i <= len(b); i++ {
		da := a[len(a)-i]
		db := b[len(b)-i]
		res[len(res)-i], rest = addDigits(da, db, rest)
	}

	for ; i <= len(a); i++ {
		res[len(res)-i], rest = addDigits(a[len(a)-i], zero, rest)
	}

	if rest != 0 {
		res[0] = digits[rest-minDigit]
	}

	if res[0] == 0 {
		res = res[1:]
	}

	return res
}

func addDigits(d1, d2 byte, modifier int) (byte, int) {
	d := dVals[d1] + dVals[d2] + modifier

	rest := 0
	if d > maxDigit {
		rest = 1
		d = d - baseSize
	}
	if d < minDigit {
		rest = -1
		d = d + baseSize
	}

	return digits[d-minDigit], rest
}
