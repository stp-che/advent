package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/10/solution"
)

func main() {
	s := sigStrenChecker{}
	solution.New(&s).Run(os.Args[1])
	fmt.Println(s.signalStrengh)
}

type sigStrenChecker struct {
	signalStrengh int
}

func (s *sigStrenChecker) Observe(cycle, x int) {
	if cycle%40 == 20 {
		s.signalStrengh += x * cycle
		// fmt.Printf("cycle: %d, x: %d, stren: %d, strenTotal: %d\n", cycle, x, x*cycle, signalStrengh)
	}
}
