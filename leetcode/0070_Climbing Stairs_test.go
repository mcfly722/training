package leetcode_test

import "testing"

/*
70. Climbing Stairs
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?



Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:
1 <= n <= 45

*/

// Fibonachi sequence
// 1,2,3(2+1),5(3+2),8(5+3)...
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	a := 3
	b := 2
	for i := 0; i < n-3; i++ {
		a = a + b
		b = a - b
	}
	return a
}

/* // recursive solution
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}

	c := 0
	for i := 1; i <= 2; i++ {
		c = c + climbStairs(n-i)
	}

	return c
}
*/

func Test_0070_Example1(t *testing.T) {
	if climbStairs(2) != 2 {
		t.Fatal()
	}
}

func Test_0070_Example2(t *testing.T) {
	if climbStairs(3) != 3 {
		t.Fatal()
	}
}
