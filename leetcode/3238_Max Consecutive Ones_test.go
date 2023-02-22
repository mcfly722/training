package main_test

import (
	"testing"
)

func findMaxConsecutiveOnes(nums []int) int {
	max := 0

	for i := 0; i < len(nums); i++ {
		j := 0
		for ; (i+j < len(nums)) && (nums[i] == nums[i+j]) && (nums[i] == 1); j++ {
		}
		if j > max {
			max = j
		}
	}

	return max
}

func Test_3238_Example1(t *testing.T) {
	if findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}) != 3 {
		t.Fatal()
	}
}
func Test_3238_Example2(t *testing.T) {
	if findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}) != 2 {
		t.Fatal()
	}
}

func Test_3238_Example3_Zero(t *testing.T) {
	if findMaxConsecutiveOnes([]int{0}) != 0 {
		t.Fatal()
	}
}
