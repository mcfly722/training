package main_test

import (
	"math"
	"testing"
)

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
