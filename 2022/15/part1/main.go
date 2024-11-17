package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/15/solution"
)

func main() {
	res := solution.New().CoverageInRow(os.Args[1], os.Args[2])
	fmt.Println(res)
}
