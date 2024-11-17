package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/9/solution"
)

func main() {
	res := solution.New(10).Run(os.Args[1])
	fmt.Println(res)
}
