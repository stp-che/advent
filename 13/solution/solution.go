package solution

import (
	"bufio"
	"log"
	"os"
)

type itemObserver interface {
	Observe(Item)
}

type Solution struct {
	observer itemObserver
}

func New(o itemObserver) *Solution {
	return &Solution{o}
}

func (s *Solution) Run(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s.handleInput(file)
}

func (s *Solution) handleInput(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		s.observer.Observe(parsePack(scanner.Text()))
	}
}

type Item interface {
	Cmp(Item) int
}

type NumItem int

func (n NumItem) Cmp(other Item) int {
	switch i := other.(type) {
	case NumItem:
		if n < i {
			return -1
		}
		if n > i {
			return 1
		}
	case ListItem:
		return -i.Cmp(n)
	}
	return 0
}

type ListItem []Item

func (l ListItem) Cmp(other Item) int {
	switch oi := other.(type) {
	case NumItem:
		return l.Cmp(ListItem{oi})
	case ListItem:
		for i := 0; i < len(oi) && i < len(l); i++ {
			res := l[i].Cmp(oi[i])
			if res != 0 {
				return res
			}
		}
		if len(l) < len(oi) {
			return -1
		}
		if len(l) > len(oi) {
			return 1
		}
	}
	return 0
}

func parsePack(str string) Item {
	var numCtx bool
	var num int
	listsStack := []ListItem{{}}

	push := func(l ListItem) {
		listsStack = append(listsStack, l)
	}

	pop := func() ListItem {
		n := len(listsStack)
		lst := listsStack[n-1]
		listsStack = listsStack[:n-1]

		return lst
	}

	appendToTop := func(i Item) {
		n := len(listsStack)
		listsStack[n-1] = append(listsStack[n-1], i)
	}

	endItem := func() {
		if numCtx {
			appendToTop(NumItem(num))
			numCtx = false
			num = 0
		}
	}

	for i := 0; i < len(str); i++ {
		switch rune(str[i]) {
		case '[':
			push(ListItem{})
		case ']':
			endItem()
			appendToTop(pop())
		case ',':
			endItem()
		case ' ':
			continue
		default:
			numCtx = true
			num = num*10 + digit(str[i])
		}
	}

	return listsStack[0][0]
}

const zeroByte = byte('0')

func digit(b byte) int {
	return int(b - zeroByte)
}
