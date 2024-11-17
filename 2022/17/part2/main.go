package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/17/solution"
)

func main() {
	// res := solution.New(1000000000000).Run(os.Args[1])
	res := solution.New(1000000000000).Run(os.Args[1])
	fmt.Println(res)
}
