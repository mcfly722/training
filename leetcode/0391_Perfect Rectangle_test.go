package main_test

import (
	"math"
	"testing"
)

/*
391. Perfect Rectangle
Given an array rectangles where rectangles[i] = [xi, yi, ai, bi] represents an axis-aligned rectangle. The bottom-left point of the rectangle is (xi, yi) and the top-right point of it is (ai, bi).

Return true if all the rectangles together form an exact cover of a rectangular region.



Example 1:


Input: rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
Output: true
Explanation: All 5 rectangles together form an exact cover of a rectangular region.
Example 2:


Input: rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
Output: false
Explanation: Because there is a gap between the two rectangular regions.
Example 3:


Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
Output: false
Explanation: Because two of the rectangles overlap with each other.


Constraints:

1 <= rectangles.length <= 2 * 104
rectangles[i].length == 4
-105 <= xi, yi, ai, bi <= 105

*/

func isTwoRectanglesAreOverlapping(rec1 []int, rec2 []int) bool {
	return rec1[0] < rec2[2] &&
		rec1[1] < rec2[3] &&
		rec2[0] < rec1[2] &&
		rec2[1] < rec1[3]
}

func isRectangleCover(rectangles [][]int) bool {
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := math.MinInt32
	maxY := math.MinInt32

	totalSquare := 0

	for i, rectangle := range rectangles {
		minX = MinInt(minX, rectangle[0])
		minX = MinInt(minX, rectangle[2])
		maxX = MaxInt(maxX, rectangle[0])
		maxX = MaxInt(maxX, rectangle[2])

		minY = MinInt(minY, rectangle[1])
		minY = MinInt(minY, rectangle[3])
		maxY = MaxInt(maxY, rectangle[1])
		maxY = MaxInt(maxY, rectangle[3])

		square := (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])

		totalSquare += square

		for j, rectangle2 := range rectangles {
			if i != j {
				if isTwoRectanglesAreOverlapping(rectangle, rectangle2) {
					return false
				}
			}
		}
		//fmt.Println(fmt.Sprintf("%v square=%v", rectangle, square))
	}

	supposedSquare := (maxX - minX) * (maxY - minY)

	//fmt.Println(fmt.Sprintf("boundaries: (%v,%v,%v,%v) supposedSquare= totalSquare=%v", minX, minY, maxX, maxY, totalSquare))

	return supposedSquare == totalSquare
}

func Test_0391_Example1(t *testing.T) {
	if !isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}) {
		t.Fatal()
	}
}

func Test_0391_Example2(t *testing.T) {
	if isRectangleCover([][]int{
		{1, 1, 2, 3},
		{1, 3, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
	}) {
		t.Fatal()
	}
}

func Test_0391_Example3(t *testing.T) {
	if isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{2, 2, 4, 4},
	}) {
		t.Fatal()
	}
}

func Test_0391_Example4(t *testing.T) {
	if isRectangleCover([][]int{
		{0, 0, 1, 1},
		{0, 1, 3, 2},
		{1, 0, 2, 2},
	}) {
		t.Fatal()
	}
}
