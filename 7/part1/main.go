package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/7/solution"
)

func main() {
	h := &sizeHandler{}
	solution.New(h, nil).Run(os.Args[1])
	fmt.Println(h.total)
}

const maxDirSize = 100000

type sizeHandler struct {
	total int
}

func (h *sizeHandler) HandleSize(size int) {
	if size <= maxDirSize {
		h.total += size
	}
}
