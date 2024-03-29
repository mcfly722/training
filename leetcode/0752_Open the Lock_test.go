package leetcode_test

import "testing"

/*
752. Open the Lock
You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.



Example 1:

Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".
Example 2:

Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".
Example 3:

Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation: We cannot reach the target without getting stuck.


Constraints:

1 <= deadends.length <= 500
deadends[i].length == 4
target.length == 4
target will not be in the list deadends.
target and deadends[i] consist of digits only.
*/

func getNeighbours(sample string) []string {
	neighbours := []string{}

	for i := 0; i < len(sample); i++ {

		current := []byte(sample)

		current[i] = ((current[i]-'0')+1)%10 + '0'
		neighbours = append(neighbours, string(current))

		current[i] = (10+(current[i]-'0')-2)%10 + '0'
		neighbours = append(neighbours, string(current))

	}

	return neighbours
}

func openLock(deadends []string, target string) int {

	visited := map[string]int{}
	queue := []string{}

	for _, deadend := range deadends {
		visited[deadend] = 0
	}

	if _, isInitialPositionInDeadends := visited["0000"]; isInitialPositionInDeadends {
		return -1
	}

	queue = append(queue, "0000")
	visited["0000"] = 0

	for len(queue) > 0 {

		// pull element from queue
		element := queue[0]
		queue = queue[1:]

		currentSteps := visited[element]

		if element == target {
			return currentSteps
		}

		for _, neighbour := range getNeighbours(element) {
			if _, alreadyVisited := visited[neighbour]; !alreadyVisited {
				queue = append(queue, neighbour)
				visited[neighbour] = currentSteps + 1
			}
		}

	}

	return -1
}

func Test_0752_Example1(t *testing.T) {
	steps := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	if steps != 6 {
		t.Fatal(steps)
	}
}

func Test_0752_Example2(t *testing.T) {
	steps := openLock([]string{"8888"}, "0009")
	if steps != 1 {
		t.Fatal(steps)
	}
}

func Test_0752_Example3(t *testing.T) {
	steps := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	if steps != -1 {
		t.Fatal(steps)
	}
}

func Test_0752_Example4(t *testing.T) {
	steps := openLock([]string{}, "0000")
	if steps != 0 {
		t.Fatal(steps)
	}
}

func Test_0752_Example5(t *testing.T) {
	steps := openLock([]string{"0000"}, "8888")
	if steps != -1 {
		t.Fatal(steps)
	}
}

func Test_0752_Example6(t *testing.T) {
	steps := openLock([]string{"1111", "0000"}, "8888")
	if steps != -1 {
		t.Fatal(steps)
	}
}
