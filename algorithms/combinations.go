package algorithms

//#include "combinations.h"
import "C"

func Multiply(a int, b int) int {
	return int(C.Multiply(C.int(a), C.int(b)))
}

func NextSequence(array *[]byte, digits byte) bool {
	if len(*array) == 0 {
		return false
	}

	if (*array)[0] < digits {
		(*array)[0]++
		return true
	}

	(*array)[0] = 0

	newArr := (*array)[1:]
	return NextSequence(&newArr, digits)
}
