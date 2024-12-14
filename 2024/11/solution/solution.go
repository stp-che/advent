package solution

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	debug bool
	cache map[string]map[int]int
}

func New(debug bool) *Solution {
	return &Solution{debug: debug, cache: map[string]map[int]int{}}
}

func (s *Solution) Run(inputPath string, blinks int) int {
	values := s.readData(inputPath)

	sum := 0
	for _, v := range values {
		sum += s.countStones(v, blinks)
	}

	return sum
}

func (s *Solution) countStones(value string, blinks int) int {
	if value == "" {
		return 0
	}
	if blinks == 0 {
		return 1
	}
	if amount, ok := s.cache[value][blinks]; ok {
		return amount
	}

	v1, v2 := s.blink(value)
	if s.cache[value] == nil {
		s.cache[value] = map[int]int{}
	}
	s.cache[value][blinks] = s.countStones(v1, blinks-1) + s.countStones(v2, blinks-1)

	return s.cache[value][blinks]
}

func (s *Solution) blink(value string) (string, string) {
	if value == "0" {
		return "1", ""
	} else if len(value)%2 == 0 {
		m := len(value) / 2
		return value[:m], withoutLeadZeros(value[m:])
	} else {
		n, _ := strconv.Atoi(value)
		return strconv.Itoa(n * 2024), ""
	}
}

func (s *Solution) readData(inputPath string) []string {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), " ")
}

func withoutLeadZeros(s string) string {
	i := 0
	for s[i] == '0' && i < len(s)-1 {
		i++
	}
	return s[i:]
}
