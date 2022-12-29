package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stp-che/advent/2022/21/solution"
)

func main() {
	start := time.Now()
	res := solution.New().Part1(os.Args[1])
	fmt.Println(res)
	fmt.Printf("time taken: %v\n", time.Since(start))
}
