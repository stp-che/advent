package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/13/solution"
)

func main() {
	o := newPackagesObserver([]solution.Item{
		solution.ListItem{solution.ListItem{solution.NumItem(2)}},
		solution.ListItem{solution.ListItem{solution.NumItem(6)}},
	})
	solution.New(o).Run(os.Args[1])
	fmt.Println(o.countersProduct())
}

type packagesObserver struct {
	dividers []solution.Item
	counters []int
}

func newPackagesObserver(dividers []solution.Item) *packagesObserver {
	return &packagesObserver{
		dividers: dividers,
		counters: make([]int, len(dividers)),
	}
}

func (o *packagesObserver) Observe(item solution.Item) {
	for i, d := range o.dividers {
		if item.Cmp(d) < 0 {
			o.counters[i]++
		}
	}
}

func (o *packagesObserver) countersProduct() int {
	res := 1
	for i, c := range o.counters {
		res *= c + i + 1
	}

	return res
}
