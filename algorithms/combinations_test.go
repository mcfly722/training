package algorithms_test

import (
	"fmt"
	"github.com/mcfly722/training/algorithms"
	"testing"
)

func Test_Multiply(t *testing.T) {
	a := 5
	b := 6
	r := algorithms.Multiply(a, b)
	if r != a*b {
		t.Fatalf(fmt.Sprintf("obtained wrong answer %v*%v=%v", a, b, r))
	}
}

func Test_NextSequence(t *testing.T) {
	array := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	i := 0
	for {
		//t.Logf(fmt.Sprintf("%5v %v", i, array))
		if !algorithms.NextSequence(&array, 3) {
			break
		}
		i++
	}
	t.Logf(fmt.Sprintf("total operations: %7v", i))

}

func Benchmark_NextSequence(b *testing.B) {

	for i := 0; i < b.N; i++ {
		array := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	operation:
		for {
			if !algorithms.NextSequence(&array, 3) {
				break operation
			}
		}
	}
}
