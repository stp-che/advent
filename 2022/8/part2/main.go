package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/8/solution"
)

func main() {
	res := solution.New(handleInput).Run(os.Args[1])
	fmt.Println(res)
}

func handleInput(data [][]byte) int {
	rows := len(data)
	cols := len(data[0])

	maxScore := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			score := scenicScore(data, i, j)

			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func scenicScore(data [][]byte, row, col int) int {
	rows := len(data)
	cols := len(data[0])
	cur := data[row][col]
	var k int

	res := 1

	for k = row - 1; k > 0 && cur > data[k][col]; k-- {
	}

	res *= row - k

	for k = row + 1; k < rows-1 && cur > data[k][col]; k++ {
	}

	res *= k - row

	for k = col - 1; k > 0 && cur > data[row][k]; k-- {
	}

	res *= col - k

	for k = col + 1; k < cols-1 && cur > data[row][k]; k++ {
	}

	res *= k - col

	return res
}
