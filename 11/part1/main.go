package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/11/solution"
)

func main() {
	res := solution.New(3, 20).Run(os.Args[1])
	fmt.Println(res)
}
