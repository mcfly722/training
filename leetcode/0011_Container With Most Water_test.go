package main_test

import (
	"testing"
)

/*
11. Container With Most Water

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104


*/

func maxArea(height []int) int {
	result := 0

	left := 0
	right := len(height) - 1

	for left < right {
		area := (right - left) * MinInt(height[left], height[right])
		//fmt.Println(fmt.Sprintf("%v[%v],%v[%v] ->min=%v, area=%v)", height[left], left, height[right], right, MinInt(height[left], height[right]), area))
		result = MaxInt(result, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func Test_0011_Example1(t *testing.T) {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		t.Fatal(result)
	}
}

func Test_0011_Example2(t *testing.T) {
	result := maxArea([]int{1, 1})
	if result != 1 {
		t.Fatal(result)
	}
}
