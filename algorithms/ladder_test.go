package algorithms_test

import (
	"fmt"
	"github.com/mcfly722/training/algorithms"
	"testing"
)

func Test_NewLadder(t *testing.T) {
	ladder := algorithms.NewLadder(5)
	t.Logf(fmt.Sprintf("%v", ladder))
}

func Test_NextLadder_Count_352715(t *testing.T) {

	ladder := algorithms.NewLadder(10)

	i := 0
	for {
		//t.Logf(fmt.Sprintf("%7v  %v", i, ladder))
		if algorithms.NextLadder(ladder, 21) == -1 {
			break
		}
		i++
	}
	if i != 352715 {
		t.Fatalf(fmt.Sprintf("total operations: %7v", i))
	}
}

func Benchmark_NextLadder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ladder := algorithms.NewLadder(10)
		for algorithms.NextLadder(ladder, 21) != -1 {
		}
	}
}

func Test_NextLadder_cgo_Count(t *testing.T) {
	i := algorithms.NextLadder_cgo_Count_352715_batch(100)
	if i != 352715*100 {
		t.Fatalf(fmt.Sprintf("total operations: %7v", i))
	}
}

func Benchmark_NextLadder_cgo(b *testing.B) {
	algorithms.NextLadder_cgo_Count_352715_batch(b.N)
}
