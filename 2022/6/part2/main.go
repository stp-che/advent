package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/6/solution"
)

func main() {
	res := solution.New(14).Run(os.Args[1])
	fmt.Println(res)
}
