package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/stp-che/advent/2024/11/solution"
)

func main() {
	start := time.Now()
	blinks, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	res := solution.New(false).Run(os.Args[1], blinks)
	fmt.Println(res)
	fmt.Printf("time taken: %v\n", time.Since(start))
}
