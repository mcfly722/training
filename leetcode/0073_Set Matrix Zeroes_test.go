package leetcode_test

import (
	"fmt"
	"testing"
)

/*
73. Set Matrix Zeroes
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.



Example 1:


Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
Example 2:


Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]


Constraints:

m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1


Follow up:

A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

func setZeroes(matrix [][]int) {
	zeroX := -1
	zeroY := -1

	// search first zero
zeroFound:
	for y, row := range matrix {
		for x := range row {
			if matrix[y][x] == 0 {
				zeroX = x
				zeroY = y
				break zeroFound
			}
		}
	}

	if zeroX != -1 && zeroY != -1 {

		//fmt.Println(fmt.Sprintf("1. (%v,%v) %v", zeroX, zeroY, matrix))

		// mark all verticals and horizontals with 0 if it should be cleaned
		for y, row := range matrix {
			for x := range row {
				if x != zeroX && y != zeroY {

					if matrix[y][x] == 0 {
						matrix[zeroY][x] = 0
						matrix[y][zeroX] = 0
					}

				}
			}
		}

		//fmt.Println(fmt.Sprintf("2: (%v,%v) %v", zeroX, zeroY, matrix))

		// clean verticals
		for x := range matrix[zeroY] {
			if matrix[zeroY][x] == 0 && x != zeroX {
				for y := 0; y < len(matrix); y++ {
					matrix[y][x] = 0
				}
			}
		}

		// clean horizontals
		for y := range matrix {
			if matrix[y][zeroX] == 0 && y != zeroY {
				for x := 0; x < len(matrix[y]); x++ {
					matrix[y][x] = 0
				}
			}
		}

		//fmt.Println(fmt.Sprintf("3: (%v,%v) %v", zeroX, zeroY, matrix))

		// clear zero horizontal
		for x := 0; x < len(matrix[zeroY]); x++ {
			matrix[zeroY][x] = 0
		}

		// clear zero vertival
		for y := 0; y < len(matrix); y++ {
			matrix[y][zeroX] = 0
		}

	}

}

func Test_0073_Example1(t *testing.T) {
	grid := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(grid)

	if fmt.Sprintf("%v", grid) != "[[1 0 1] [0 0 0] [1 0 1]]" {
		t.Fatal(fmt.Sprintf("%v", grid))
	}
}

func Test_0073_Example2(t *testing.T) {
	grid := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(grid)

	if fmt.Sprintf("%v", grid) != "[[0 0 0 0] [0 4 5 0] [0 3 1 0]]" {
		t.Fatal(fmt.Sprintf("%v", grid))
	}
}

func Test_0073_Example3(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 1, 0},
	}

	setZeroes(grid)

	if fmt.Sprintf("%v", grid) != "[[1 0 1 0] [0 0 0 0] [0 0 0 0]]" {
		t.Fatal(fmt.Sprintf("%v", grid))
	}
}
