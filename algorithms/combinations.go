package algorithms

/*
int Multiply(int a, int b) {
	return a * b;

}
*/
import "C"

func Multiply(a int, b int) int {
	return C.Multiply(C.int(a), C.int(b))
}
