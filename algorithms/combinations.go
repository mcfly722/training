package algorithms

//#include "combinations.h"
import "C"

func Multiply(a int, b int) int {
	return int(C.Multiply(C.int(a), C.int(b)))
}
