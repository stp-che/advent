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

	visibleInside := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if visible(data, i, j) {
				visibleInside += 1
			}
		}
	}

	return (rows+cols-2)*2 + visibleInside
}

func visible(data [][]byte, row, col int) bool {
	rows := len(data)
	cols := len(data[0])
	cur := data[row][col]
	var k int

	for k = row - 1; k >= 0 && cur > data[k][col]; k-- {
	}

	if k < 0 {
		return true
	}

	for k = row + 1; k < rows && cur > data[k][col]; k++ {
	}

	if k == rows {
		return true
	}

	for k = col - 1; k >= 0 && cur > data[row][k]; k-- {
	}

	if k < 0 {
		return true
	}

	for k = col + 1; k < cols && cur > data[row][k]; k++ {
	}

	if k == cols {
		return true
	}

	return false
}
