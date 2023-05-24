package algorithms

//#include "combinations.h"
import "C"

func Multiply(a int, b int) int {
	return int(C.Multiply(C.int(a), C.int(b)))
}

func NextSequenceFastest(array *[]byte, digits byte) bool {

	for i := 0; (*array)[i] < digits; i++ {
		(*array)[i]++

		if (*array)[i] < digits {
			return true
		}
		(*array)[i] = 0
	}

	return false
}

func NextSequenceRecursive(array *[]byte, digits byte) bool {

	if len(*array) == 0 {
		return false
	}

	if (*array)[0] < digits-1 {
		(*array)[0]++
		return true
	}

	(*array)[0] = 0

	newArr := (*array)[1:]
	return NextSequenceRecursive(&newArr, digits)
}
