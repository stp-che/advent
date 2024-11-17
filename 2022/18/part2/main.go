package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/18/solution"
)

func main() {
	res := solution.New().Run1(os.Args[1])
	fmt.Println(res)
}
