package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/10/solution"
)

func main() {
	solution.New(&display{}).Run(os.Args[1])
}

type display struct{}

func (d *display) Observe(cycle, x int) {
	diff := (cycle-1)%40 - x
	c := "."
	if -2 < diff && diff < 2 {
		c = "#"
	}
	fmt.Print(c)
	if cycle%40 == 0 {
		fmt.Println()
	}
}
