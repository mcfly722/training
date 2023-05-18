package algorithms_test

import (
	"algorithms"
	"fmt"
	"testing"
)

func Test_algo_0001(t *testing.T) {
	r := algorithms.Multiply(5, 6)
	fmt.Println(fmt.Sprintf("%v", r))
}
