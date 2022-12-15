package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) CoverageInRow(inputPath string, row string) int {
	return s.prepare(inputPath).coverageInRow(parseInt(row))
}

func (s *Solution) InspectRegion(inputPath string, bound string) (int, int) {
	return s.prepare(inputPath).inspectRegion(parseInt(bound))
}

func (s *Solution) prepare(inputPath string) scanning {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sensors := make([]sensor, 0)

	for scanner.Scan() {
		sensors = append(sensors, parseSensor(scanner.Text()))
	}

	return scanning{sensors}
}

const sensorPattern = "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d"

func parseSensor(s string) sensor {
	var sx, sy, bx, by int
	fmt.Sscanf(s, sensorPattern, &sx, &sy, &bx, &by)

	pos := point{sx, sy}
	bpos := point{bx, by}

	return sensor{
		pos:    pos,
		bpos:   bpos,
		radius: distance(pos, bpos),
	}
}

type scanning struct {
	sensors []sensor
}

func (s scanning) coverageInRow(y int) int {
	u := union{}
	bs := map[int]bool{}

	for _, sensor := range s.sensors {
		present, coverage := sensor.coverageInRow(y)

		if !present {
			continue
		}

		u = u.add(coverage)

		if sensor.bpos.y == y {
			bs[sensor.bpos.x] = true
		}
	}

	return u.size() - len(bs)
}

func (s scanning) inspectRegion(bound int) (int, int) {
	for y := 0; y <= bound; y++ {
		u := union{{0, bound}}

		for _, sensor := range s.sensors {
			present, coverage := sensor.coverageInRow(y)

			if !present {
				continue
			}

			u = u.remove(coverage)
		}

		if len(u) > 0 {
			return u[0].b, y
		}
	}

	return 0, 0
}

type interval struct {
	b, e int
}

func newInterval(b, e int) interval {
	if b > e {
		log.Fatalf("interval begin (%d) is greater than end (%d)\n", b, e)
	}

	return interval{b, e}
}

func (i interval) length() int {
	return i.e - i.b + 1
}

type point struct {
	x, y int
}

type sensor struct {
	pos    point
	bpos   point
	radius int
}

func (s sensor) coverageInRow(y int) (bool, interval) {
	dx := s.radius - abs(s.pos.y-y)
	if dx < 0 {
		return false, newInterval(0, 0)
	}

	return true, newInterval(s.pos.x-dx, s.pos.x+dx)
}

func distance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func parseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(n)
}

type union []interval

func (u union) add(in interval) union {
	if len(u) == 0 {
		return union{in}
	}

	if u[0].b > in.e+1 {
		return union(append(union{in}, u...))
	}

	if u[len(u)-1].e < in.b-1 {
		return union(append(u, in))
	}

	right := -1
	for i := 0; i < len(u); i++ {
		if in.e <= u[i].e {
			right = i
			break
		}
	}

	left := -1
	for i := len(u) - 1; i >= 0; i-- {
		if in.b >= u[i].b {
			left = i
			break
		}
	}

	// fmt.Printf("left: %d, right: %d\n", left, right)

	if left == -1 && right == -1 {
		return union{in}
	}

	leftPart := union{}
	rightPart := union{}
	mb, me := in.b, in.e

	if left >= 0 {
		i := left + 1
		if u[left].e+1 >= in.b {
			i--
			mb = u[left].b
		}
		leftPart = u[:i]
	}

	if right >= 0 {
		i := right
		if u[right].b-1 <= in.e {
			i++
			me = u[right].e
		}
		rightPart = u[i:]
	}

	// fmt.Println(leftPart)
	// fmt.Println(newInterval(mb, me))
	// fmt.Println(rightPart)

	return join(leftPart, union{newInterval(mb, me)}, rightPart)
}

func (u union) remove(in interval) union {
	if len(u) == 0 || u[0].b > in.e || u[len(u)-1].e < in.b {
		return u
	}

	left := 0
	for ; u[left].e < in.b; left++ {
	}
	leftPart := u[:left]

	right := len(u) - 1
	for ; u[right].b > in.e; right-- {

	}
	rightPart := u[right+1:]

	middle := union{}

	if u[left].b < in.b {
		middle = append(middle, newInterval(u[left].b, in.b-1))
	}

	if u[right].e > in.e {
		middle = append(middle, newInterval(in.e+1, u[right].e))
	}

	// fmt.Println(leftPart)
	// fmt.Println(middle)
	// fmt.Println(rightPart)

	return join(leftPart, middle, rightPart)
}

func (u union) size() int {
	size := 0
	for _, i := range u {
		size += i.length()
	}

	return size
}

func join(us ...union) union {
	n := 0
	for _, u := range us {
		n += len(u)
	}

	new := make(union, n)
	i := 0
	for _, u := range us {
		for j := 0; j < len(u); j++ {
			new[i] = u[j]
			i++
		}
	}

	return new
}
