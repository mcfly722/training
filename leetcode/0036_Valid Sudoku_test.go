package leetcode_test

import "testing"

/*
36. Valid Sudoku
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.


Example 1:


Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.


Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

func isValidSudoku(board [][]byte) bool {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

	// check horizontals
	for y := 0; y < 9; y++ {
		primesLine := 1
		for x := 0; x < 9; x++ {
			if board[y][x] != '.' {
				value := board[y][x] - '1'
				if (primesLine)%(primes[value]) == 0 {
					return false
				} else {
					primesLine *= primes[value]
				}
			}
		}
	}

	// check verticals
	for x := 0; x < 9; x++ {
		primesLine := 1
		for y := 0; y < 9; y++ {
			if board[y][x] != '.' {
				value := board[y][x] - '1'
				if (primesLine)%(primes[value]) == 0 {
					return false
				} else {
					primesLine *= primes[value]
				}
			}
		}
	}

	// check squares
	for xS := 0; xS < 3; xS++ {
		for yS := 0; yS < 3; yS++ {
			primesLine := 1
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[yS*3+y][xS*3+x] != '.' {
						value := board[yS*3+y][xS*3+x] - '1'
						if (primesLine)%(primes[value]) == 0 {
							return false
						} else {
							primesLine *= primes[value]
						}
					}
				}
			}
		}
	}

	return true
}

func Test_0036_Example1(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	if !isValidSudoku(board) {
		t.Fatal()
	}
}

func Test_0036_Example2(t *testing.T) {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	if isValidSudoku(board) {
		t.Fatal()
	}
}

func Test_0036_Example3(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '.'},
		{'.', '6', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '8', '.'},
		{'.', '.', '.', '.', '3', '.', '.', '.', '8'},
	}

	if isValidSudoku(board) {
		t.Fatal()
	}
}

func Test_0036_Example4(t *testing.T) {
	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}

	if isValidSudoku(board) {
		t.Fatal()
	}
}
