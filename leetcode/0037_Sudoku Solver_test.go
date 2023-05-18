package leetcode_test

import (
	"math"
	"testing"
)

/*
37. Sudoku Solver
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
The '.' character indicates empty cells.



Example 1:


Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
Explanation: The input board is shown above and the only valid solution is shown below:




Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit or '.'.
It is guaranteed that the input board has only one solution.
*/

func digitsDoesNotContainDupes(digits []byte) bool {
	var dupes [10]int

	dupes = [len(dupes)]int{}

	for _, digit := range digits {
		if digit != '.' {
			i := (digit - '1') % 10
			if dupes[i] == 0 {
				dupes[i] = 1
			} else {
				return false
			}
		}
	}

	return true
}

func IsCorrectVertical(board [][]byte, column int) bool {
	digits := []byte{}

	for y := 0; y < 9; y++ {
		c := board[y][column]
		if c != '.' {
			digits = append(digits, c)
		}
	}

	return digitsDoesNotContainDupes(digits)
}

func IsCorrectHorizontal(board [][]byte, row int) bool {
	digits := []byte{}

	for x := 0; x < 9; x++ {
		c := board[row][x]
		if c != '.' {
			digits = append(digits, c)
		}
	}

	return digitsDoesNotContainDupes(digits)
}

func IsCorrectSmallSquare(board [][]byte, column int, row int) bool {
	xs := (int)(math.Floor(float64(column / 3)))
	ys := (int)(math.Floor(float64(row / 3)))

	digits := []byte{}

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			c := board[ys*3+y][xs*3+x]
			if c != '.' {
				digits = append(digits, c)
			}
		}
	}

	return digitsDoesNotContainDupes(digits)
}

func IsCompleted(board [][]byte) bool {
	for y, row := range board {
		for x := range row {
			if board[y][x] == '.' {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]byte) {

	if IsCompleted(board) {
		//fmt.Println(TwoDimensionalByteArray2String(board))
		return
	}

	for y, row := range board {
		for x := range row {

			if board[y][x] == '.' {

				for v := 1; v <= 9; v++ {

					board[y][x] = '0' + byte(v)

					if IsCorrectVertical(board, x) && IsCorrectHorizontal(board, y) && IsCorrectSmallSquare(board, x, y) {
						solveSudoku(board)
						if IsCompleted(board) {
							return
						}
					}
				}

				board[y][x] = '.'

				return
			}
		}
	}
}

func Test_0037_Example1(t *testing.T) {

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

	solveSudoku(board)

	solution := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}

	for y, row := range board {
		for x := range row {
			if board[y][x] != solution[y][x] {
				t.Fatal(TwoDimensionalByteArray2String(board))
			}
		}
	}
}
