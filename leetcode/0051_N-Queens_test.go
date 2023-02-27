package main_test

import (
	"fmt"
	"testing"
)

/*
51. N-Queens
Hard
9.5K
211
Companies
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.



Example 1:


Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
Example 2:

Input: n = 1
Output: [["Q"]]


Constraints:

1 <= n <= 9
*/

type queensField struct {
	grid        [][]byte
	diagsXY     []byte
	diagsYX     []byte
	horizontals []byte
	verticals   []byte
	size        int
	count       int
}

func newQueensField(n int) *queensField {
	queensField := &queensField{
		grid:        [][]byte{},
		diagsXY:     []byte{},
		diagsYX:     []byte{},
		horizontals: []byte{},
		verticals:   []byte{},
		size:        n,
	}

	for y := 0; y < n; y++ {
		row := []byte{}
		for x := 0; x < n; x++ {
			row = append(row, '.')
		}
		queensField.grid = append(queensField.grid, row)
	}

	for i := 0; i < n; i++ {
		queensField.horizontals = append(queensField.horizontals, 0)
		queensField.verticals = append(queensField.verticals, 0)
	}

	for i := 0; i < n*2; i++ {
		queensField.diagsXY = append(queensField.diagsXY, 0)
		queensField.diagsYX = append(queensField.diagsYX, 0)
	}

	return queensField
}

func (queensField *queensField) Put(x, y int) bool {
	if queensField.horizontals[y] == 0 && queensField.verticals[x] == 0 && queensField.diagsXY[x+y] == 0 && queensField.diagsYX[queensField.size-x+y] == 0 {
		queensField.horizontals[y] = 1
		queensField.verticals[x] = 1
		queensField.diagsXY[x+y] = 1
		queensField.diagsYX[queensField.size-x+y] = 1
		queensField.count++
		queensField.grid[y][x] = 'Q'
		return true
	}
	return false
}

func (queensField *queensField) Pop(x, y int) {
	if queensField.grid[y][x] != '.' {
		queensField.horizontals[y] = 0
		queensField.verticals[x] = 0
		queensField.diagsXY[x+y] = 0
		queensField.diagsYX[queensField.size-x+y] = 0
		queensField.grid[y][x] = '.'
		queensField.count--
	}
}

func (queensField *queensField) bfs(onSolutionFound func(field *queensField)) {

	if queensField.count == queensField.size {
		onSolutionFound(queensField)
	}

	for y := 0; y < queensField.size; y++ {
		if y <= queensField.count { // to do not skip the lines
			for x := 0; x < queensField.size; x++ {
				if queensField.Put(x, y) {
					queensField.bfs(onSolutionFound)
					queensField.Pop(x, y)
				}
			}
		}
	}
}

func (queensField *queensField) toStringsArray() []string {
	solution := []string{}
	for y := 0; y < queensField.size; y++ {
		row := ""
		for x := 0; x < queensField.size; x++ {
			row += string(queensField.grid[y][x])
		}
		solution = append(solution, row)
	}
	return solution
}

func solveNQueens(n int) [][]string {
	field := newQueensField(n)

	solutions := [][]string{}

	onSolutionFound := func(field *queensField) {
		solutions = append(solutions, field.toStringsArray())
	}

	field.bfs(onSolutionFound)

	return solutions
}

func Test_0051_Example1(t *testing.T) {
	result := fmt.Sprintf("%v", solveNQueens(1))
	requiredAnswer := "[[Q]]"
	if result != requiredAnswer {
		t.Fatal(result)
	}
}

func Test_0051_Example2(t *testing.T) {
	result := fmt.Sprintf("%v", solveNQueens(2))
	requiredAnswer := "[]"
	if result != requiredAnswer {
		t.Fatal(result)
	}
}
func Test_0051_Example3(t *testing.T) {
	result := fmt.Sprintf("%v", solveNQueens(3))
	requiredAnswer := "[]"
	if result != requiredAnswer {
		t.Fatal(result)
	}
}
func Test_0051_Example4(t *testing.T) {
	result := fmt.Sprintf("%v", solveNQueens(4))
	requiredAnswer := "[[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]"
	if result != requiredAnswer {
		t.Fatal(result)
	}
}

func Test_0051_Example5(t *testing.T) {
	result := fmt.Sprintf("%v", solveNQueens(5))
	requiredAnswer := "[[Q.... ..Q.. ....Q .Q... ...Q.] [Q.... ...Q. .Q... ....Q ..Q..] [.Q... ...Q. Q.... ..Q.. ....Q] [.Q... ....Q ..Q.. Q.... ...Q.] [..Q.. Q.... ...Q. .Q... ....Q] [..Q.. ....Q .Q... ...Q. Q....] [...Q. Q.... ..Q.. ....Q .Q...] [...Q. .Q... ....Q ..Q.. Q....] [....Q .Q... ...Q. Q.... ..Q..] [....Q ..Q.. Q.... ...Q. .Q...]]"
	if result != requiredAnswer {
		t.Fatal(result)
	}
}
