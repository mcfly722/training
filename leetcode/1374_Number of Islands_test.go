package main_test

import (
	"fmt"
	"testing"
)

/*
Number of Islands

Solution
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.



Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

func dfsDeleteLand(x int, y int, grid [][]byte) {

	if grid[y][x] != '0' {
		grid[y][x] = '0'
		if x > 0 {
			dfsDeleteLand(x-1, y, grid)
		}

		if x < len(grid[y])-1 {
			dfsDeleteLand(x+1, y, grid)
		}

		if y > 0 {
			dfsDeleteLand(x, y-1, grid)
		}

		if y < len(grid)-1 {
			dfsDeleteLand(x, y+1, grid)
		}
	}
}

func numIslands(grid [][]byte) int {
	count := 0

	for y, row := range grid {
		for x := range row {
			if grid[y][x] != '0' {
				count++

				dfsDeleteLand(x, y, grid)

				//				fmt.Println(fmt.Sprintf("%v\n", TwoDimensionalByteArray2String(grid)))

			}
		}
	}

	return count
}

func Test_1374_Example1(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	result := numIslands(grid)
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_1374_Example2(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	result := numIslands(grid)

	fmt.Println(result)

	if result != 3 {
		t.Fatal(result)
	}
}
