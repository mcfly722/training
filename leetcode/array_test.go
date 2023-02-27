package main_test

import (
	"testing"
)

func IsEqualIntArrays(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func IsEqualIntMatricies(first, second [][]int) bool {
	if len(first) != len(second) {
		return false
	}
	for y, row := range first {
		if len(row) != len(second[y]) {
			return false
		}
		for x := range row {
			if first[y][x] != second[y][x] {
				return false
			}

		}
	}
	return true
}

func IsEqualStringsArrays(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for y, s := range first {
		if s != second[y] {
			return false
		}
	}
	return true
}

func IsEqualByteMatricies(first, second [][]byte) bool {
	if len(first) != len(second) {
		return false
	}
	for y, row := range first {
		if len(row) != len(second[y]) {
			return false
		}
		for x := range row {
			if first[y][x] != second[y][x] {
				return false
			}

		}
	}
	return true
}

func IsEqualByteArrays(first, second []byte) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func IsEqualStringArrays(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func SearchInsertPosInSortedIntArrayWithRight(nums *[]int, target int, right int) int {
	if len(*nums) == 0 {
		return 0
	}
	if target <= (*nums)[0] {
		return 0
	}
	if target > (*nums)[len(*nums)-1] {
		return len(*nums)
	}
	left := 0
	for {
		p := (left + right) / 2
		if target > (*nums)[p] {
			left = p
		}
		if target <= (*nums)[p] {
			right = p
		}
		if left == right-1 {
			return right
		}
	}
}

func SearchInsertPosInSortedIntArray(nums *[]int, target int) int {
	return SearchInsertPosInSortedIntArrayWithRight(nums, target, len(*nums)-1)
}

func InstertIntoSortedIntArray(array *[]int, target int) {
	pos := SearchInsertPosInSortedIntArray(array, target)
	if pos == len(*array) {
		*array = append(*array, target)
		return
	}

	*array = append((*array)[:pos+1], (*array)[pos:]...)
	(*array)[pos] = target
}

func RemoveFromSortedIntArray(array *[]int, target int) {
	pos := SearchInsertPosInSortedIntArray(array, target)
	if pos < len(*array) {
		*array = append((*array)[:pos], (*array)[pos+1:]...)
	}
}

func RemoveDuplicatesInSortedIntArray(nums *[]int) {

	if len(*nums) < 2 {
		return
	}

	ic := 0 // current sorted index

	for i := 1; i < len(*nums); i++ {

		if (*nums)[i] > (*nums)[ic] {
			ic++
			(*nums)[ic] = (*nums)[i]
		}
	}

	(*nums) = (*nums)[:ic+1]
}

func TwoDimensionalByteArray2String(array [][]byte) string {
	result := ""

	for _, row := range array {
		for _, cell := range row {
			result += string(cell)
		}
		result += "\n"
	}

	return result
}

func GetIntArrayMaximum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func GetIntArrayMinimum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Test_IsEqualIntArrays_Example1(t *testing.T) {
	if !IsEqualIntArrays([]int{}, []int{}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example2(t *testing.T) {
	if !IsEqualIntArrays([]int{1}, []int{1}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example3(t *testing.T) {
	if IsEqualIntArrays([]int{1}, []int{2}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example4(t *testing.T) {
	if !IsEqualIntArrays([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fatal()
	}
}

func Test_IsEqualIntArrays_Example5(t *testing.T) {
	if IsEqualIntArrays([]int{1, 2, 5}, []int{1, 2, 3}) {
		t.Fatal()
	}
}
