package main_test

import (
	"math"
	"testing"
)

/*
149. Max Points on a Line
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, return the maximum number of points that lie on the same straight line.

Example 1:

Input: points = [[1,1],[2,2],[3,3]]
Output: 3
Example 2:

Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4


Constraints:

1 <= points.length <= 300
points[i].length == 2
-104 <= xi, yi <= 104
All the points are unique.

*/

func maxPoints(points [][]int) int {

	max := 0

	// single point
	if len(points)==1 {
		return 1
	}

	for i := 0; i < len(points); i++ {

		slopesCounter := map[float64]int{}

		for j := 0; j < len(points); j++ {

			if i != j {
				slope := math.MaxFloat64

				if (points[i])[0]-(points[j])[0] != 0 {
					slope = (float64)(points[i][1]-points[j][1]) / (float64)(points[i][0]-points[j][0])
				}

				if _, found := slopesCounter[slope]; found {
					slopesCounter[slope]++
				} else {
					slopesCounter[slope] = 2
				}

				//				fmt.Println(fmt.Sprintf("%v,%v", points[i], points[j]))

			}
		}

		for _, slopesCounter := range slopesCounter {
			max = MaxInt(max, slopesCounter)
		}

		//		fmt.Println(fmt.Sprintf("       %v max=%v", len(slopesCounter), max))

	}

	return max
}

func Test_0149_Example1(t *testing.T) {
	sample := [][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}

	if maxPoints(sample) != 4 {
		t.Fatal()
	}
}

func Test_0149_Example2(t *testing.T) {
	sample := [][]int{}

	if maxPoints(sample) != 0 {
		t.Fatal()
	}
}

func Test_0149_Example3(t *testing.T) {
	sample := [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{0, 1},
	}

	if maxPoints(sample) != 3 {
		t.Fatal()
	}
}

func Test_0149_Example4(t *testing.T) {
	sample := [][]int{
		{1, 1},
		{2, 1},
		{3, 1},
		{0, 2},
	}

	if maxPoints(sample) != 3 {
		t.Fatal()
	}
}

func Test_0149_Example5(t *testing.T) {
	sample := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	if maxPoints(sample) != 3 {
		t.Fatal()
	}
}

func Test_0149_Example5(t *testing.T) {
	sample := [][]int{
		{0, 0}
	}

	if maxPoints(sample) != 1 {
		t.Fatal()
	}
}
