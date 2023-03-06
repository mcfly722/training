package main_test

import (
	"fmt"
	"testing"
)

func CantorPairingFunction(x, y int) int64 {
	return (int64)((x+y)*(x+y+1)/2 + y)
}

func CantorPairingFunctionWithNegative(x, y int) int64 {
	x = x * 2
	y = y * 2
	if x < 0 {
		x = -1*x + 1
	}
	if y < 0 {
		y = -1*y + 1
	}
	return CantorPairingFunction(x, y)
}

func CantorPairingFunctionWithNegativeAndCommutative(x, y int) int64 {
	max := MaxInt(x, y)
	min := MinInt(x, y)

	return CantorPairingFunctionWithNegative(min, max)
}

func CantorPairingFunctionWithNegativeAndCommutativeForThree(x, y, z int) int64 {

	max := MaxInt(x, MaxInt(y, z))
	min := MinInt(x, MinInt(y, z))
	mid := x + y + z - max - min

	return CantorPairingFunctionWithNegative((int)(CantorPairingFunctionWithNegative(min, mid)), max)
}

type pair struct {
	x int
	y int
}

func Test_PairingCommutativeNegative_Range(t *testing.T) {
	results := map[int64]pair{}

	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			unique := CantorPairingFunctionWithNegativeAndCommutative(x, y)

			if p, found := results[unique]; found {
				if p.x+p.y != x+y { // not commutative sum

					t.Fatal(fmt.Sprintf("duplicate found :(%v,%v)=(%v,%v)=%v", x, y, p.x, p.y, unique))
				}
			} else {
				results[unique] = pair{x: x, y: y}
			}
		}
		//fmt.Println(x)
	}
}

func Test_PairingCommutativeNegative_Commutation(t *testing.T) {
	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			if CantorPairingFunctionWithNegativeAndCommutative(x, y) != CantorPairingFunctionWithNegativeAndCommutative(y, x) {
				t.Fatal(fmt.Sprintf("point %v,%v is not commutative!", x, y))
			}
		}
		//fmt.Println(fmt.Sprintf("%v", x))
	}
}
