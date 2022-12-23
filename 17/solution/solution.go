package solution

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stp-che/advent/2022/17/solution/alt"
)

const (
	width = 7
)

type Solution struct {
	steps int
}

func New(steps int) *Solution {
	return &Solution{steps}
}

func (s *Solution) Run(inputPath string) int {
	winds, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	// fGen := newCycleFormsGenerator([][]int{
	// 	formHorLine, formPlus, formAngle, formVertLine, formBox,
	// })
	fGen := newCycleFiguressGenerator(width, [][]byte{
		formHorLine, formPlus, formAngle, formVertLine, formBox,
	})

	sim := newSimulation(width, 3, 2, winds, fGen)
	// sim = sim.verbose()

	start := time.Now()

	for i := 0; i < s.steps; i++ {
		sim.dropNextFigure()
		if i%10000 == 0 {
			x := float64(i+1) / float64(s.steps)
			fmt.Printf("\r%0.2f%%, eta: %s", x*100, time.Now().Add(time.Duration(float64(time.Since(start))/x)))
		}
	}
	fmt.Println()

	return sim.stackHeight()
}

func (s *Solution) Run1(inputPath string) int {
	winds, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	fGen := alt.NewCycleFormsGenerator([]alt.Form{
		alt.FormHorLine, alt.FormPlus, alt.FormAngle, alt.FormVertLine, alt.FormBox,
	})

	sim := alt.NewSimulation(7, 3, 2, winds, fGen)
	sim = sim.Verbose()

	start := time.Now()

	for i := 0; i < s.steps; i++ {
		sim.DropNextFigure()
		if i%10000 == 0 {
			x := float64(i+1) / float64(s.steps)
			fmt.Printf("\r%0.2f%%, eta: %s", x*100, time.Now().Add(time.Duration(float64(time.Since(start))/x)))
		}
	}
	fmt.Println()

	return sim.StackHeight()
}
