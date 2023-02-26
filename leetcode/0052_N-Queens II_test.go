package main_test

import (
	"testing"
)

func totalNQueens(n int) int {
	return len(solveNQueens(n)) //solveNQueens declared in 0051
}

func Test_0052_Example1(t *testing.T) {
	result := totalNQueens(1)
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0052_Example4(t *testing.T) {
	result := totalNQueens(4)
	if result != 2 {
		t.Fatal(result)
	}
}

func Test_0052_Example5(t *testing.T) {
	result := totalNQueens(5)
	if result != 10 {
		t.Fatal(result)
	}
}

func Test_0052_Example6(t *testing.T) {
	result := totalNQueens(6)
	if result != 4 {
		t.Fatal(result)
	}
}

func Test_0052_Example7(t *testing.T) {
	result := totalNQueens(7)
	if result != 40 {
		t.Fatal(result)
	}
}

func Test_0052_Example8(t *testing.T) {
	result := totalNQueens(8)
	if result != 92 {
		t.Fatal(result)
	}
}
