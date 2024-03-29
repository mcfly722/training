package leetcode_test

import (
	"fmt"
	"math"
	"testing"
)

/*
8. String to Integer (atoi)
Medium
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

Read in and ignore any leading whitespace.
Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
Return the integer as the final result.
Note:

Only the space character ' ' is considered a whitespace character.
Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.


Example 1:

Input: s = "42"
Output: 42
Explanation: The underlined characters are what is read in, the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
The parsed integer is 42.
Since 42 is in the range [-231, 231 - 1], the final result is 42.
Example 2:

Input: s = "   -42"
Output: -42
Explanation:
Step 1: "   -42" (leading whitespace is read and ignored)
            ^
Step 2: "   -42" ('-' is read, so the result should be negative)
             ^
Step 3: "   -42" ("42" is read in)
               ^
The parsed integer is -42.
Since -42 is in the range [-231, 231 - 1], the final result is -42.
Example 3:

Input: s = "4193 with words"
Output: 4193
Explanation:
Step 1: "4193 with words" (no characters read because there is no leading whitespace)
         ^
Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
             ^
The parsed integer is 4193.
Since 4193 is in the range [-231, 231 - 1], the final result is 4193.


Constraints:

0 <= s.length <= 200
s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.

*/

func myAtoi(s string) int {
	num := int64(0)
	sign := int64(1)

	i := 0

	if len(s) == 0 {
		return 0
	}

	// search end of spaces
	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}

	// if just spaces
	if i > len(s)-1 {
		return 0
	}

	//fmt.Println(fmt.Sprintf("spaces finished at:%v", i))

	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	if i > len(s)-1 || !(s[i] >= '0' && s[i] <= '9') {
		return 0
	}

	//fmt.Println(fmt.Sprintf("sign:%v", sign))

	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		num = num*10 + (int64)(s[i]-'0')

		if num > math.MaxInt32 {
			if sign == 1 {
				num = math.MaxInt32
			} else {
				num = math.MaxInt32 + 1
			}
		}

	}

	//	fmt.Println(fmt.Sprintf("%v", num*sign))

	return int(num * sign)
}

func Test_0008_Example1(t *testing.T) {
	if myAtoi("  123  4") != 123 {
		t.Fatal()
	}
}

func Test_0008_Example2(t *testing.T) {
	if myAtoi("   0  ") != 0 {
		t.Fatal()
	}
}

func Test_0008_Example3(t *testing.T) {
	if myAtoi("  -123  ") != -123 {
		t.Fatal()
	}
}

func Test_0008_Example4(t *testing.T) {
	if myAtoi("-123 ") != -123 {
		t.Fatal()
	}
}

func Test_0008_Example5(t *testing.T) {
	if myAtoi("-123 ") != -123 {
		t.Fatal()
	}
}

func Test_0008_Example6(t *testing.T) {
	if myAtoi(fmt.Sprintf(" %v", math.MaxInt32+5)) != math.MaxInt32 {
		t.Fatal()
	}
}

func Test_0008_Example7(t *testing.T) {
	if myAtoi(fmt.Sprintf(" %v", -(math.MaxInt32+5))) != -(math.MaxInt32 + 1) {
		t.Fatal()
	}
}

func Test_0008_Example8(t *testing.T) {
	if myAtoi("-9") != -9 {
		t.Fatal()
	}
}

func Test_0008_Example9(t *testing.T) {
	if myAtoi("") != 0 {
		t.Fatal()
	}
}

func Test_0008_Example10(t *testing.T) {
	if myAtoi("-") != 0 {
		t.Fatal()
	}
}

func Test_0008_Example11(t *testing.T) {
	if myAtoi("21474836460") != 2147483647 {
		t.Fatal()
	}
}

func Test_0008_Example12(t *testing.T) {
	if myAtoi("-91283472332") != -2147483648 {
		t.Fatal()
	}
}

func Test_0008_Example13(t *testing.T) {
	if myAtoi(" ") != 0 {
		t.Fatal()
	}
}

func Test_0008_Example14(t *testing.T) {
	if myAtoi("1") != 1 {
		t.Fatal()
	}
}
