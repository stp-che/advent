package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	handleInput(input)
}

func handleInput(input []byte) {
	withOverlap := 0
	for _, row := range strings.Split(string(input), "\n") {
		if haveOverlap(parseRow(row)) {
			withOverlap++
		}
	}
	fmt.Println(withOverlap)
}

func parseRow(row string) ([2]int, [2]int) {
	ranges := strings.Split(row, ",")
	return parseRange(ranges[0]), parseRange(ranges[1])
}

func parseRange(rStr string) [2]int {
	r := strings.Split(rStr, "-")
	return [2]int{parseInt(r[0]), parseInt(r[1])}
}

func parseInt(s string) int {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(res)
}

func haveOverlap(r1, r2 [2]int) bool {
	if r1[0] > r2[0] {
		r1, r2 = r2, r1
	}

	return r1[1] >= r2[0]
}
