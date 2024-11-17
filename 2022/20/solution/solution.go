package solution

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	mixTimes   int
	multipiler int
}

func New(mixTimes, multiplier int) *Solution {
	return &Solution{mixTimes, multiplier}
}

func (s *Solution) Run(inputPath string) int {
	items, zeroItem := s.readData(inputPath)
	n := len(items)

	// printItems := func() {
	// 	i := zeroItem
	// 	for j := 0; j < n; j++ {
	// 		fmt.Print(i.v, " ")
	// 		i = i.next
	// 	}
	// 	fmt.Println()
	// }

	for i := 0; i < s.mixTimes; i++ {

		for _, item := range items {
			// printItems()

			sh := item.v % (n - 1)

			if sh > 0 {
				moveForward(item, sh)
				continue
			}

			moveBackward(item, -sh)
		}
		// printItems()
	}

	a := findFrom(zeroItem, 1000)
	b := findFrom(zeroItem, 2000)
	c := findFrom(zeroItem, 3000)

	fmt.Println(a.v)
	fmt.Println(b.v)
	fmt.Println(c.v)

	return a.v + b.v + c.v
}

func moveBackward(i *item, n int) {
	if n == 0 {
		return
	}

	prevItem := i.prev
	for j := 0; j < n; j++ {
		prevItem = prevItem.prev
	}

	moveItemAfter(i, prevItem)
}

func moveForward(i *item, n int) {
	if n == 0 {
		return
	}

	prevItem := i
	for j := 0; j < n; j++ {
		prevItem = prevItem.next
	}

	moveItemAfter(i, prevItem)
}

func moveItemAfter(i, prevItem *item) {
	i.prev.next = i.next
	i.next.prev = i.prev

	i.prev = prevItem
	i.next = prevItem.next
	prevItem.next.prev = i
	prevItem.next = i
}

func findFrom(i *item, n int) *item {
	res := i
	for j := 0; j < n; j++ {
		res = res.next
	}

	return res
}

func (s *Solution) readData(inputPath string) ([]*item, *item) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(data), "\n")
	items := make([]*item, len(rows))
	var prev *item
	var zeroItem *item

	for i, row := range rows {
		v, err := strconv.ParseInt(row, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		items[i] = &item{v: int(v) * s.multipiler, prev: prev}
		if prev != nil {
			prev.next = items[i]
		}

		prev = items[i]

		if v == 0 {
			zeroItem = items[i]
		}
	}

	items[len(items)-1].next = items[0]
	items[0].prev = items[len(items)-1]

	return items, zeroItem
}

type item struct {
	v    int
	prev *item
	next *item
}
