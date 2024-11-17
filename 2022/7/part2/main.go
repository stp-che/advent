package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/7/solution"
)

func main() {
	f := &finder{
		totalSpace:    70000000,
		requiredSpace: 30000000,
	}
	solution.New(f, f).Run(os.Args[1])
	fmt.Println(f.removeCandidateSize)
}

type finder struct {
	totalSpace          int
	requiredSpace       int
	usedSpace           int
	removeCandidateSize int
}

func (f *finder) SetSpaceUsed(size int) {
	f.usedSpace = size
}

func (f *finder) HandleSize(size int) {
	if size >= f.lackOfSpace() && (size < f.removeCandidateSize || f.removeCandidateSize == 0) {
		f.removeCandidateSize = size
	}
}

func (f *finder) lackOfSpace() int {
	return f.requiredSpace - (f.totalSpace - f.usedSpace)
}
