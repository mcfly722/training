package main_test

import (
	"fmt"
	"testing"
)

/*
412. Fizz Buzz
Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.


Example 1:

Input: n = 3
Output: ["1","2","Fizz"]
Example 2:

Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
Example 3:

Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]


Constraints:

1 <= n <= 104
*/

func fizzBuzz(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, fmt.Sprintf("%v", i))
		}
	}

	return result
}

func Test_0412_Example1(t *testing.T) {
	result := fizzBuzz(3)
	if !stringArraysIsEqual(result, []string{"1", "2", "Fizz"}) {
		t.Fatal(result)
	}
}

func Test_0412_Example2(t *testing.T) {
	result := fizzBuzz(5)
	if !stringArraysIsEqual(result, []string{"1", "2", "Fizz", "4", "Buzz"}) {
		t.Fatal(result)
	}
}

func Test_0412_Example3(t *testing.T) {
	result := fizzBuzz(15)
	if !stringArraysIsEqual(result, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}) {
		t.Fatal(result)
	}
}
