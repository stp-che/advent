package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stp-che/advent/2022/22/solution"
)

func main() {
	start := time.Now()
	res := solution.New(50).Run(os.Args[1])
	fmt.Println(res)
	fmt.Printf("time taken: %v\n", time.Since(start))
}
