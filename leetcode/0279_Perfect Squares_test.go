package main_test

import (
	"fmt"
	"math"
	"testing"
)

/*
Perfect Squares

Solution
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.



Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.


Constraints:

1 <= n <= 104
*/

func dfsShortestPath(currentPath []int, shortestPath *[]int, leavedPart int) {

	if len(currentPath) >= len(*shortestPath) {
		return
	}

	if leavedPart == 0 {
		*shortestPath = currentPath
		return
	}

	if leavedPart < 0 {
		return
	}

	for i := int(math.Sqrt(float64(leavedPart))); i >= 1; i-- {

		newPath := append(currentPath, i*i)

		dfsShortestPath(newPath, shortestPath, leavedPart-i*i)

	}

}

func numSquares(n int) int {

	if n == 0 {
		return 1
	}

	shortestPath := []int{}
	for i := 0; i < n; i++ {
		shortestPath = append(shortestPath, 1)
	}

	dfsShortestPath([]int{}, &shortestPath, n)

	fmt.Println(fmt.Sprintf("%v", shortestPath))
	return len(shortestPath)
}

func Test_0279_Example1(t *testing.T) {
	result := numSquares(12)
	if result != 3 {
		t.Fatal(result)
	}
}

func Test_0279_Example2(t *testing.T) {
	result := numSquares(13)
	if result != 2 {
		t.Fatal(result)
	}
}

func Test_0279_Example3(t *testing.T) {
	result := numSquares(0)
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0279_Example4(t *testing.T) {
	result := numSquares(1)
	if result != 1 {
		t.Fatal(result)
	}
}
