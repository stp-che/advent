package solution

import (
	"bufio"
	"log"
	"os"
)

type Solution struct {
	markerSize int
}

func New(markerSize int) *Solution {
	return &Solution{markerSize}
}

func (s *Solution) Run(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return s.handleInput(file)
}

func (s *Solution) handleInput(file *os.File) int {
	var res = -1

	eachByte(file, bytesHandler(s.markerSize, &res))

	return res + 1
}

func bytesHandler(n int, offset *int) func(int, byte) bool {
	uniqSeq := make([]byte, 0, n)

	return func(i int, b byte) bool {
		uniqSeq = append(uniqSeq, b)

		for j := len(uniqSeq) - 2; j >= 0; j-- {
			if uniqSeq[j] == b {
				uniqSeq = uniqSeq[j+1:]
				break
			}
		}

		// fmt.Println(string(uniqSeq))

		if len(uniqSeq) < n {
			return true
		}

		*offset = i
		return false
	}
}

func eachByte(file *os.File, handler func(int, byte) bool) {
	reader := bufio.NewReader(file)
	var b byte
	var err error
	var offset = 0
	for {
		b, err = reader.ReadByte()
		if err != nil {
			break
		}

		if !handler(offset, b) {
			break
		}

		offset++
	}
}
