package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/14/solution"
)

func main() {
	res := solution.New(2).Run(os.Args[1])
	fmt.Println(res)
}
