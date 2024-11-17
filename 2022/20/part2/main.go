package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stp-che/advent/2022/20/solution"
)

func main() {
	start := time.Now()
	res := solution.New(10, 811589153).Run(os.Args[1])
	fmt.Println(res)
	fmt.Printf("time taken: %v\n", time.Since(start))
}
