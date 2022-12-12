package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/12/solution"
)

func main() {
	res := solution.New().ShortestPathFromAnyA(os.Args[1])
	fmt.Println(res)
}
