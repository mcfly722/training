package algorithms

//#include "combinations.h"
import "C"

func Multiply(a int, b int) int {
	return int(C.Multiply(C.int(a), C.int(b)))
}

func NextSequenceFastest(array *[]byte, base byte) bool {

	for i := 0; (*array)[i] < base; i++ {
		(*array)[i]++

		if (*array)[i] < base {
			return true
		}
		(*array)[i] = 0
	}

	return false
}

func NextSequenceRecursive(array *[]byte, base byte) bool {

	if len(*array) == 0 {
		return false
	}

	if (*array)[0] < base-1 {
		(*array)[0]++
		return true
	}

	(*array)[0] = 0

	newArr := (*array)[1:]
	return NextSequenceRecursive(&newArr, base)
}

func NextSequenceFastest_cgo_Count_1048575_batch(n int) uint64 {
	return uint64(C.NextSequenceFastest_Count_1048575_batch(C.int(n)))
}
