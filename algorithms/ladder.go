package algorithms

//#include "ladder.h"
import "C"

/*
	ladder(height) -> gaps -> sequence

	123(3) -> 000 -> 11100
	124(4) -> 001 -> 11010
	125(5) -> 002 -> 11001
	134(4) -> 010 -> 10110
	135(5) -> 011 -> 10101
	145(5) -> 020 -> 10011
	234(4) -> 100 -> 01110
	235(5) -> 101 -> 01101
	245(5) -> 110 -> 01011
	345(5) -> 200 -> 00111

*/

/*
	NewLadder initiates array with size=n. Each element are greater than previous.
	NewLadder(5) = [1,2,3,4,5]
	NewLadder(2) = [1,2]
	etc.
*/
func NewLadder(n int) *[]int {
	array := []int{}
	for i := 0; i < n; i++ {
		array = append(array, i+1)
	}
	return &array
}

/*
	NextLadder - fast generation of next ladder which is less than maxHeight
	returns last modification position in ladder
*/
func NextLadder(array *[]int, maxHeight int) int {

	// get last correct modification position
	length := len(*array)
	pos := length - 1
	max := maxHeight - pos

	for ; pos > -1 && (*array)[pos] >= max+pos; pos-- {
	}

	if pos == -1 {
		return -1
	}

	r := pos

	// increment it
	(*array)[pos]++
	pos++

	// reset tail
	for ; pos < length; pos++ {
		(*array)[pos] = (*array)[pos-1] + 1
	}

	return r
}

func NextLadder_cgo_Count_352715_batch(n int) uint64 {
	return uint64(C.NextLadder_Count_352715_batch(C.int(n)))
}
