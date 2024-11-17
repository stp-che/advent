package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/11/solution"
)

func main() {
	res := solution.New(1, 10000).Run(os.Args[1])
	fmt.Println(res)
}
