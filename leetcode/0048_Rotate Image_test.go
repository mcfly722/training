package main_test

import (
	"fmt"
	"testing"
)

/*
48. Rotate Image
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.



Example 1:


Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
Example 2:


Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]


Constraints:

n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/

func transponate(matrix [][]int) {
	rows := len(matrix)
	cols := rows

	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

}

func horizontalMirror(matrix [][]int) {
	for y, row := range matrix {
		for x := 0; x < len(row)/2; x++ {
			tmp := matrix[y][x]
			matrix[y][x] = matrix[y][len(row)-1-x]
			matrix[y][len(row)-1-x] = tmp
		}
	}
}

func rotateMatrix(matrix [][]int) {
	transponate(matrix)
	horizontalMirror(matrix)
}

func Test_0770_Example1(t *testing.T) {
	sample := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	requiredAnswer := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	rotateMatrix(sample)

	if !IsEqualIntMatricies(sample, requiredAnswer) {
		t.Fatal(fmt.Sprintf("%v", sample))
	}
}

func Test_0770_Example2(t *testing.T) {
	sample := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	requiredAnswer := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	rotateMatrix(sample)

	if !IsEqualIntMatricies(sample, requiredAnswer) {
		t.Fatal(fmt.Sprintf("%v", sample))
	}
}

func Test_0770_Example3(t *testing.T) {
	sample := [][]int{
		{4},
	}
	requiredAnswer := [][]int{{4}}

	rotateMatrix(sample)

	if !IsEqualIntMatricies(sample, requiredAnswer) {
		t.Fatal(fmt.Sprintf("%v", sample))
	}
}

func Test_0770_Example4(t *testing.T) {
	sample := [][]int{}
	requiredAnswer := [][]int{}

	rotateMatrix(sample)

	if !IsEqualIntMatricies(sample, requiredAnswer) {
		t.Fatal(fmt.Sprintf("%v", sample))
	}
}
