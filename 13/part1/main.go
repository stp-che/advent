package main

import (
	"fmt"
	"os"

	"github.com/stp-che/advent/2022/13/solution"
)

func main() {
	o := newPairsObserver()
	solution.New(o).Run(os.Args[1])
	fmt.Println(o.rightOrderIndicesSum)
}

type pairsObserver struct {
	pair                 [2]solution.Item
	packageIndex         int
	rightOrderIndicesSum int
}

func newPairsObserver() *pairsObserver {
	return &pairsObserver{
		pair: [2]solution.Item{},
	}
}

func (o *pairsObserver) Observe(item solution.Item) {
	i := o.packageIndex % 2
	o.pair[i] = item
	if i == 1 {
		fmt.Printf("%v\n%v\n%d\n\n", o.pair[0], o.pair[1], o.pair[0].Cmp(o.pair[1]))
		if o.pair[0].Cmp(o.pair[1]) < 1 {
			o.rightOrderIndicesSum += o.packageIndex/2 + 1
		}
	}

	o.packageIndex++
}
