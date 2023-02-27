package main_test

import (
	"math"
	"testing"
)

/*
204. Count Primes
Given an integer n, return the number of prime numbers that are strictly less than n.

Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
Example 2:

Input: n = 0
Output: 0
Example 3:

Input: n = 1
Output: 0


Constraints:

0 <= n <= 5 * 106
*/

func countPrimes(n int) int {

	if n < 3 {
		return 0
	}

	primes := []int{2}

	for i := 3; i < n; i = i + 2 {

		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
			if p > (int)(math.Sqrt((float64)(i))) {
				break
			}
		}

		if isPrime == true {
			primes = append(primes, i)
		}

	}

	return len(primes)
}

func Test_0204_Example1(t *testing.T) {
	result := countPrimes(10)
	if result != 4 {
		t.Fatal(result)
	}
}

func Test_0204_Example2(t *testing.T) {
	result := countPrimes(0)
	if result != 0 {
		t.Fatal(result)
	}
}
func Test_0204_Example3(t *testing.T) {
	result := countPrimes(1)
	if result != 0 {
		t.Fatal(result)
	}
}

func Test_0204_Example4(t *testing.T) {
	result := countPrimes(2)
	if result != 0 {
		t.Fatal(result)
	}
}

func Test_0204_Example5(t *testing.T) {
	result := countPrimes(3)
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0204_Example6(t *testing.T) {
	result := countPrimes(499979)
	if result != 41537 {
		t.Fatal(result)
	}
}
