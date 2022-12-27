package solution

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Part1(inputPath string, t int) int {
	bs := s.readData(inputPath)

	out := make([]chan int, len(bs))

	for i := 0; i < len(bs); i++ {
		out[i] = make(chan int)
		go func(out chan<- int, b *blueprint) {
			out <- b.qLevel(t)
		}(out[i], bs[i])
	}

	sum := 0
	for _, ch := range out {
		sum += <-ch
	}

	return sum
}

func (s *Solution) Part2(inputPath string, t int) int {
	bs := s.readData(inputPath)[:3]

	out := make([]chan int, len(bs))

	for i := 0; i < len(bs); i++ {
		out[i] = make(chan int)
		go func(out chan<- int, b *blueprint) {
			out <- b.maxGeods(t)
		}(out[i], bs[i])
	}

	prod := 1
	for _, ch := range out {
		prod *= <-ch
	}

	return prod
}

func (s *Solution) readData(inputPath string) []*blueprint {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	res := make([]*blueprint, 0)

	for _, s := range strings.Split(string(data), "\n") {
		res = append(res, parseBlueprint1(s))
	}

	return res
}

type resVal [4]int

func (v resVal) includes(other resVal) bool {
	for i, n := range v {
		if n < other[i] {
			return false
		}
	}

	return true
}

func (v resVal) add(other resVal) resVal {
	res := resVal{}

	for i := 0; i < len(v); i++ {
		res[i] = v[i] + other[i]
	}

	return res
}

func (v resVal) sub(other resVal) resVal {
	res := resVal{}

	for i := 0; i < len(v); i++ {
		res[i] = v[i] - other[i]
	}

	return res
}

func (v resVal) mult(n int) resVal {
	res := resVal{}

	for i := 0; i < len(v); i++ {
		res[i] = v[i] * n
	}

	return res
}

const (
	resOre   = 0
	resClay  = 1
	resObsid = 2
	resGeod  = 3
)

type blueprint struct {
	id                int
	prodIncreaseCosts [4]resVal
}

func (b *blueprint) qLevel(t int) int {
	return b.id * b.maxGeods(t)
}

func (b *blueprint) maxGeods(t int) int {
	fmt.Println(b)

	var rec func(int, resVal, resVal) int

	// pad := "                                                                                    "
	// i := -1

	rec = func(t int, curProd, stock resVal) int {
		// i++
		// defer func() {
		// 	i--
		// }()

		// fmt.Println(string(pad[:i]), t, curProd, stock)

		if t < 2 {
			return stock[resGeod] + t*curProd[resGeod]
		}

		if !superMax(t-2, curProd, stock).includes(b.prodIncreaseCosts[resGeod]) {
			return stock[resGeod] + t*curProd[resGeod]
		}

		max := t * curProd[resGeod]

		for i := 0; i < len(curProd); i++ {
			takenTime, nextStock := nextTo(stock, b.prodIncreaseCosts[i], curProd)

			if takenTime == -1 || takenTime >= t {
				continue
			}

			curProd[i]++
			v := rec(t-takenTime, curProd, nextStock)
			curProd[i]--

			if max < v {
				max = v
			}
		}

		return max
	}

	return rec(t, resVal{1, 0, 0, 0}, resVal{})
}

func nextTo(stock, stockTarget, prod resVal) (int, resVal) {
	max := 0
	val := resVal{}

	for i := 0; i < len(stockTarget); i++ {
		t := minsTo(stock[i], stockTarget[i], prod[i])

		if t == -1 {
			return -1, val
		}

		if max < t {
			max = t
		}
	}

	for i := 0; i < len(val); i++ {
		val[i] = stock[i] + prod[i]*max - stockTarget[i]
	}

	return max, val
}

func minsTo(cur, target, prod int) int {
	if target <= cur {
		return 1
	}

	if prod == 0 {
		return -1
	}

	return (target-cur+prod-1)/prod + 1
}

func superMax(t int, curProd, stock resVal) resVal {
	res := resVal{}

	tt := t * (t - 1) / 2

	for i := 0; i < len(res); i++ {
		res[i] = stock[i] + t*curProd[i] + tt
	}

	return res
}

const blueprintPattern = "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian."

func parseBlueprint1(s string) *blueprint {
	var id, oror, clor, obor, obcl, geor, geob int

	fmt.Sscanf(s, blueprintPattern, &id, &oror, &clor, &obor, &obcl, &geor, &geob)

	b := blueprint{
		id: id,
		prodIncreaseCosts: [4]resVal{
			{oror, 0, 0, 0},
			{clor, 0, 0, 0},
			{obor, obcl, 0, 0},
			{geor, 0, geob, 0},
		},
	}

	// fmt.Println(b)

	return &b
}
