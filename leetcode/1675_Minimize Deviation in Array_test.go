package leetcode_test

import (
	"math"
	"testing"
)

/*

1675. Minimize Deviation in Array
You are given an array nums of n positive integers.

You can perform two types of operations on any element of the array any number of times:

If the element is even, divide it by 2.
For example, if the array is [1,2,3,4], then you can do this operation on the last element, and the array will be [1,2,3,2].
If the element is odd, multiply it by 2.
For example, if the array is [1,2,3,4], then you can do this operation on the first element, and the array will be [2,2,3,4].
The deviation of the array is the maximum difference between any two elements in the array.

Return the minimum deviation the array can have after performing some number of operations.

YOU CANNOT MULTIPLY ODD BY 2. IT IS FORBIDDEN!

Example 1:

Input: nums = [1,2,3,4]
Output: 1
Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.
Example 2:

Input: nums = [4,1,5,20,3]
Output: 3
Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.
Example 3:

Input: nums = [2,10,8]
Output: 3


Constraints:

n == nums.length
2 <= n <= 5 * 104
1 <= nums[i] <= 109
*/

func instertIntoSortedIntArrayWithoutDupes(array *[]int, target int) {

	pos := SearchInsertPosInSortedIntArray(array, target)
	if pos == len(*array) {
		*array = append(*array, target)
		return
	}

	if (*array)[pos] == target {
		return
	}

	*array = append((*array)[:pos+1], (*array)[pos:]...)
	(*array)[pos] = target
}

func minimumDeviation(nums []int) int {

	sorted := []int{}

	for _, num := range nums {
		if num%2 == 1 {
			num = num * 2
		}

		instertIntoSortedIntArrayWithoutDupes(&sorted, num)
	}

	diff := math.MaxInt32

	for {
		min := sorted[0]
		max := sorted[len(sorted)-1]

		diff = (int)(math.Min((float64)(diff), (float64)(max-min)))
		if max%2 == 0 {
			sorted = sorted[:len(sorted)-1]
			instertIntoSortedIntArrayWithoutDupes(&sorted, max/2)
		} else {
			break
		}

	}

	return diff
}

func Test_1675_Example1(t *testing.T) {
	result := minimumDeviation([]int{1, 2, 3, 4})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_1675_Example2(t *testing.T) {
	result := minimumDeviation([]int{4, 1, 5, 20, 3})
	if result != 3 {
		t.Fatal(result)
	}
}

func Test_1675_Example3(t *testing.T) {
	result := minimumDeviation([]int{2, 10, 8})
	if result != 3 {
		t.Fatal(result)
	}
}

func Test_1675_Example4(t *testing.T) {
	result := minimumDeviation([]int{5, 5, 10, 20})
	if result != 0 {
		t.Fatal(result)
	}
}

func Test_1675_Example5(t *testing.T) {
	result := minimumDeviation([]int{3, 5})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_1675_Example6(t *testing.T) {
	result := minimumDeviation([]int{4, 9, 4, 5})
	if result != 5 {
		t.Fatal(result)
	}
}
