package main_test

import (
	"sort"
	"testing"
)

/*
A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.

You are given an integer array heights representing the current order that the students are standing in. Each heights[i] is the height of the ith student in line (0-indexed).

Return the number of indices where heights[i] != expected[i].



Example 1:

Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation:
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.
Example 2:

Input: heights = [5,1,2,3,4]
Output: 5
Explanation:
heights:  [5,1,2,3,4]
expected: [1,2,3,4,5]
All indices do not match.
Example 3:

Input: heights = [1,2,3,4,5]
Output: 0
Explanation:
heights:  [1,2,3,4,5]
expected: [1,2,3,4,5]
All indices match.


Constraints:

1 <= heights.length <= 100
1 <= heights[i] <= 100
   Hide Hint #1
Build the correct order of heights by sorting another array, then compare the two arrays.
*/

func heightChecker(heights []int) int {
	differenciesCount := 0
	sortedHeights := append([]int(nil), heights...)

	sort.Slice(sortedHeights, func(i, j int) bool {
		return sortedHeights[i] < sortedHeights[j]
	})

	for i, height := range heights {
		if height != sortedHeights[i] {
			differenciesCount++
		}
	}
	return differenciesCount
}

func Test_3228_Example1(t *testing.T) {
	nums := []int{1, 1, 4, 2, 1, 3}

	diff := heightChecker(nums)

	if diff != 3 {
		t.Fatal(diff)
	}
}

func Test_3228_Example2(t *testing.T) {
	nums := []int{5, 1, 2, 3, 4}

	diff := heightChecker(nums)

	if diff != 5 {
		t.Fatal(diff)
	}
}

func Test_3228_Example3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	diff := heightChecker(nums)

	if diff != 0 {
		t.Fatal(diff)
	}
}
