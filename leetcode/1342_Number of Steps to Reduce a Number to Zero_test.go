package main_test

import "testing"

/*
1342. Number of Steps to Reduce a Number to Zero
Given an integer num, return the number of steps to reduce it to zero.

In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.



Example 1:

Input: num = 14
Output: 6
Explanation:
Step 1) 14 is even; divide by 2 and obtain 7.
Step 2) 7 is odd; subtract 1 and obtain 6.
Step 3) 6 is even; divide by 2 and obtain 3.
Step 4) 3 is odd; subtract 1 and obtain 2.
Step 5) 2 is even; divide by 2 and obtain 1.
Step 6) 1 is odd; subtract 1 and obtain 0.
Example 2:

Input: num = 8
Output: 4
Explanation:
Step 1) 8 is even; divide by 2 and obtain 4.
Step 2) 4 is even; divide by 2 and obtain 2.
Step 3) 2 is even; divide by 2 and obtain 1.
Step 4) 1 is odd; subtract 1 and obtain 0.
Example 3:

Input: num = 123
Output: 12


Constraints:

0 <= num <= 106

*/

func numberOfSteps(num int) int {
	steps := 0
	for ; num > 0; steps++ {

		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}

	}

	return steps
}

func Test_1342_Example1(t *testing.T) {
	result := numberOfSteps(14)
	if result != 6 {
		t.Fatal(result)
	}
}

func Test_1342_Example2(t *testing.T) {
	result := numberOfSteps(8)
	if result != 4 {
		t.Fatal(result)
	}
}

func Test_1342_Example3(t *testing.T) {
	result := numberOfSteps(123)
	if result != 12 {
		t.Fatal(result)
	}
}

func Test_1342_Example4(t *testing.T) {
	result := numberOfSteps(0)
	if result != 0 {
		t.Fatal(result)
	}
}

func Test_1342_Example5(t *testing.T) {
	result := numberOfSteps(1)
	if result != 1 {
		t.Fatal(result)
	}
}
