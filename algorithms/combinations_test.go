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

func Test_NextSequenceRecursive_ShowSeveral(t *testing.T) {
	array := []byte{0, 0, 0, 0, 5}
	for i := 0; ; i++ {
		t.Logf(fmt.Sprintf("%7v %v", i, array))
		if !algorithms.NextSequenceRecursive(&array, 3) {
			break
		}
	}
}

func Test_NextSequenceRecursive_Count(t *testing.T) {
	array := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5}

	i := 0
	for {
		if !algorithms.NextSequenceRecursive(&array, 4) {
			break
		}
		i++
	}
	t.Logf(fmt.Sprintf("total operations: %7v", i))

}

func Benchmark_NextSequenceRecursive(b *testing.B) {

	for i := 0; i < b.N; i++ {
		array := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5}

	operation:
		for {
			if !algorithms.NextSequenceRecursive(&array, 4) {
				break operation
			}
		}
	}
}

func Test_NextSequenceFastest_ShowSeveral(t *testing.T) {
	array := []byte{0, 0, 0, 0, 5}
	for i := 0; ; i++ {
		t.Logf(fmt.Sprintf("%7v %v", i, array))
		if !algorithms.NextSequenceFastest(&array, 3) {
			break
		}
	}
}

func Test_NextSequenceFastest_Count(t *testing.T) {
	array := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5}

	i := 0
	for {
		if !algorithms.NextSequenceFastest(&array, 4) {
			break
		}
		i++
	}
	t.Logf(fmt.Sprintf("total operations: %7v", i))

}

func Benchmark_NextSequenceFastest(b *testing.B) {

	for i := 0; i < b.N; i++ {
		array := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5}

	operation:
		for {
			if !algorithms.NextSequenceFastest(&array, 4) {
				break operation
			}
		}
	}
}
