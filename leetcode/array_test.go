package main_test

import (
	"testing"
)

func intArraysIsEqual(first, second []int) bool {
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

func byteArraysIsEqual(first, second []byte) bool {
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

func Test_intArraysIsEqual_Example1(t *testing.T) {
	if !intArraysIsEqual([]int{}, []int{}) {
		t.Fatal()
	}
}

func Test_intArraysIsEqual_Example2(t *testing.T) {
	if !intArraysIsEqual([]int{1}, []int{1}) {
		t.Fatal()
	}
}

func Test_intArraysIsEqual_Example3(t *testing.T) {
	if intArraysIsEqual([]int{1}, []int{2}) {
		t.Fatal()
	}
}

func Test_intArraysIsEqual_Example4(t *testing.T) {
	if !intArraysIsEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fatal()
	}
}

func Test_intArraysIsEqual_Example5(t *testing.T) {
	if intArraysIsEqual([]int{1, 2, 5}, []int{1, 2, 3}) {
		t.Fatal()
	}
}
