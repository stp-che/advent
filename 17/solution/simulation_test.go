package solution

import (
	"testing"
)

var winds = []byte(">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>")

// var fGen = newCycleFormsGenerator([][]int{
// 	formHorLine, formPlus, formAngle, formVertLine, formBox,
// })
var fGen = newCycleFiguressGenerator(width, [][]byte{
	formHorLine, formPlus, formAngle, formVertLine, formBox,
})

var sim = newSimulation(width, 3, 2, winds, fGen)

// $ go test -cpuprofile ../../tmp/cpu17.prof -memprofile ../../tmp/mem17.prof -bench . -benchmem
// $ go tool pprof -http=:8081 solution.test ../../tmp/cpu17.prof
// $ go tool pprof -http=:8082 solution.test ../../tmp/mem17.prof

func BenchmarkDropNextFigure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sim.dropNextFigure()
	}
}

// var fGen1 = alt.NewCycleFormsGenerator([]alt.Form{
// 	alt.FormHorLine, alt.FormPlus, alt.FormAngle, alt.FormVertLine, alt.FormBox,
// })

// var sim1 = alt.NewSimulation(7, 3, 2, winds, fGen1)

// func BenchmarkAltDropNextFigure(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		sim1.DropNextFigure()
// 	}
// }

func TestValue(t *testing.T) {
	for i := 0; i < 2022; i++ {
		sim.dropNextFigure()
		// sim1.DropNextFigure()
	}

	res := sim.stackHeight()
	expected := 3068
	// res1 := sim1.StackHeight()

	if res != expected {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", res, expected)
	}
}
