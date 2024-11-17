package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/15/solution"
)

func main() {
	x, y := solution.New().InspectRegion(os.Args[1], os.Args[2])
	fmt.Println(x*4000000 + y)
}
