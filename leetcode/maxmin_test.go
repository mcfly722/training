package main_test

import "math"

func MaxInt(a, b int) int {
	return (a + b) + (int)(math.Abs((float64)(a-b)))/2
}

func MinInt(a, b int) int {
	return (a + b) - (int)(math.Abs((float64)(a-b)))/2
}

func MaxInt64(a, b int64) int64 {
	return (a + b) + (int64)(math.Abs((float64)(a-b)))/2
}

func MinInt64(a, b int64) int64 {
	return (a + b) - (int64)(math.Abs((float64)(a-b)))/2
}
