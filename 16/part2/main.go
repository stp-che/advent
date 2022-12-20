package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/16/solution"
)

func main() {
	res := solution.New().Run(os.Args[1], 26, 2)
	fmt.Println(res)
}
