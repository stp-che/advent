package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/5/solution"
)

func main() {
	res := solution.New(moveCrates).Run(os.Args[1])
	fmt.Println(res)
}

func moveCrates(n int, from, to []byte) ([]byte, []byte) {
	newStack := make([]byte, n+len(to))
	for i := 0; i < n; i++ {
		newStack[n-1-i] = from[i]
	}
	for i := 0; i < len(to); i++ {
		newStack[n+i] = to[i]
	}
	return from[n:], newStack
}
