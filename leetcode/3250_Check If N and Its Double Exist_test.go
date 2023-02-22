package main_test

/*
Check If N and Its Double Exist
Given an array arr of integers, check if there exist two indices i and j such that :
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
Example 1:
Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
Example 2:
Input: arr = [3,1,7,11]
Output: false
Explanation: There is no i and j that satisfy the conditions.
Constraints:
2 <= arr.length <= 500
-103 <= arr[i] <= 103
*/

import "testing"

func checkIfExist(arr []int) bool {
	index := make(map[int]int)

	for n, element := range arr {
		index[element] = n
	}

	for i1, element := range arr {
		if i2, found := index[element*2]; found {
			if i1 != i2 {
				return true
			}
		}
	}

	return false
}

func Test_Found(t *testing.T) {
	if !checkIfExist([]int{10, 2, 5, 3, 10}) {
		t.Fatal()
	}
}

func Test_NotFound(t *testing.T) {
	if checkIfExist([]int{3, 1, 7, 11}) {
		t.Fatal()
	}
}

func Test_Zero(t *testing.T) {
	if checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}) {
		t.Fatal()
	}
}

func Test_ZeroAndZero(t *testing.T) {
	if !checkIfExist([]int{0, 0}) {
		t.Fatal()
	}
}
