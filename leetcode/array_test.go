package main_test

import (
	"testing"
)

func IsEqualIntArrays(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func IsEqualByteArrays(first, second []byte) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func IsEqualStringArrays(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func TwoDimensionalByteArray2String(array [][]byte) string {
	result := ""

	for _, row := range array {
		for _, cell := range row {
			result += string(cell)
		}
		result += "\n"
	}

	return result
}

func Test_IsEqualIntArrays_Example1(t *testing.T) {
	if !IsEqualIntArrays([]int{}, []int{}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example2(t *testing.T) {
	if !IsEqualIntArrays([]int{1}, []int{1}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example3(t *testing.T) {
	if IsEqualIntArrays([]int{1}, []int{2}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example4(t *testing.T) {
	if !IsEqualIntArrays([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example5(t *testing.T) {
	if IsEqualIntArrays([]int{1, 2, 5}, []int{1, 2, 3}) {
		t.Fatal()
	}
}
