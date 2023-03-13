package main_test

import (
	"fmt"
	"testing"
)

/*
60. Permutation Sequence
The set [1, 2, 3, ..., n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.



Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
Example 3:

Input: n = 3, k = 1
Output: "123"


Constraints:

1 <= n <= 9
1 <= k <= n!

*/

func copyWithoutElement_0060(slice []int, i int) []int {
	new := []int{}
	new = append(new, slice[:i]...)
	new = append(new, slice[i+1:]...)
	return new
}

func intArrayToString_0060(nums []int) string {
	result := ""

	for _, num := range nums {
		result += fmt.Sprintf("%v", num)
	}

	return result
}

func recursivePermutation(initialSet []int, path []int, leftSet []int, onFound func(solution []int) bool) bool {

	if len(leftSet) == 0 {
		needToContinue := onFound(path)
		return needToContinue
	}

	for i := 0; i < len(leftSet); i++ {
		firstWouldBe := leftSet[i]
		newLeftSet := copyWithoutElement_0060(leftSet, i)
		newPath := append(path, firstWouldBe)

		needToContinue := recursivePermutation(initialSet, newPath, newLeftSet, onFound)

		if !needToContinue {
			return needToContinue
		}
	}

	// need to continue
	return true
}

func getPermutation(n int, k int) string {

	strSolution := ""

	initialSet := []int{}

	for i := 0; i < n; i++ {
		initialSet = append(initialSet, i+1)
	}

	counter := 0

	onFound := func(solution []int) bool {
		counter++
		if counter == k {
			strSolution = intArrayToString_0060(solution)
			return false
		} else {
			return true
		}
	}

	newInitialSet := make([]int, len(initialSet))
	copy(newInitialSet, initialSet)

	recursivePermutation(initialSet, []int{}, newInitialSet, onFound)

	return strSolution
}

func Test_0060_Example1(t *testing.T) {
	result := getPermutation(3, 3)
	if result != "213" {
		t.Fatal(result)
	}
}

func Test_0060_Example2(t *testing.T) {
	result := getPermutation(4, 9)
	if result != "2314" {
		t.Fatal(result)
	}
}

func Test_0060_Example3(t *testing.T) {
	result := getPermutation(3, 1)
	if result != "123" {
		t.Fatal(result)
	}
}
