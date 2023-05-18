package algorithms_test

import (
	"fmt"
	"github.com/mcfly722/training/algorithms"
	"testing"
)

func Test_algo_0001(t *testing.T) {
	r := algorithms.Multiply(5, 6)
	t.Logf(fmt.Sprintf("%v", r))
}
