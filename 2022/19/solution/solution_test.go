package solution

import (
	"testing"
)

// $ go test -cpuprofile ../../tmp/cpu19.prof -memprofile ../../tmp/mem19.prof -bench . -benchmem
// $ go tool pprof -http=:8081 solution.test ../../tmp/cpu19.prof
// $ go tool pprof -http=:8082 solution.test ../../tmp/mem19.prof

func BenchmarkDropNextFigure(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Run("../test_input.txt", 20)
	}
}

func TestValue(t *testing.T) {
	res := New().Run("../test_input.txt", 20)
	expected := 6

	if res != expected {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", res, expected)
	}
}
