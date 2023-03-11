package main_test

import (
	"fmt"
	"testing"
)

/*
22. Generate Parentheses
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.



Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]


Constraints:

1 <= n <= 8

*/

type Step struct {
	opens  int
	closes int
}

func (step *Step) ToString() string {
	result := ""
	for opens := 0; opens < step.opens; opens++ {
		result += "("
	}
	for closes := 0; closes < step.closes; closes++ {
		result += ")"
	}
	return result
}

func recursiveGenerateParenthesis(opensLeft int, closesLeft int, steps []Step, onFound func(steps []Step)) {

	if opensLeft == 0 && closesLeft == 0 {
		onFound(steps)
	}

	for opens := 1; opens <= opensLeft; opens++ {
		for closes := 1; closes <= closesLeft; closes++ {
			if opensLeft <= closesLeft {
				newSteps := append(steps, Step{opens: opens, closes: closes})
				recursiveGenerateParenthesis(opensLeft-opens, closesLeft-closes, newSteps, onFound)
			}
		}
	}

}

func generateParenthesis(n int) []string {

	parenthesis := []string{}

	onFound := func(steps []Step) {
		solution := ""

		for _, step := range steps {
			solution += step.ToString()
		}
		parenthesis = append(parenthesis, solution)
	}

	recursiveGenerateParenthesis(n, n, []Step{}, onFound)

	return parenthesis
}

func Test_0022_Example1(t *testing.T) {
	result := generateParenthesis(4)
	if fmt.Sprintf("%v", result) != "[()()() ()(()) (()()) (())() ((()))]" {
		t.Fatal(result)
	}
}

func Test_0022_Example2(t *testing.T) {
	result := generateParenthesis(1)
	if fmt.Sprintf("%v", result) != "[()]" {
		t.Fatal(result)
	}
}
