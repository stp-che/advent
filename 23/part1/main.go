package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stp-che/advent/2022/23/solution"
)

func main() {
	start := time.Now()
	res := solution.New().RunNRounds(os.Args[1], 10)
	fmt.Println(res)
	fmt.Printf("time taken: %v\n", time.Since(start))
}
