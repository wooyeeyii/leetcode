package array

/**
1888. Minimum Number of Flips to Make the Binary String Alternating

You are given a binary string s. You are allowed to perform two types of operations on the string in any sequence:

Type-1: Remove the character at the start of the string s and append it to the end of the string.
Type-2: Pick any character in s and flip its value, i.e., if its value is '0' it becomes '1' and vice-versa.
Return the minimum number of type-2 operations you need to perform such that s becomes alternating.

The string is called alternating if no two adjacent characters are equal.

For example, the strings "010" and "1010" are alternating, while the string "0100" is not.

Example 1:
Input: s = "111000"
Output: 2
Explanation: Use the first operation two times to make s = "100011".
Then, use the second operation on the third and sixth elements to make s = "101010".

Example 2:
Input: s = "010"
Output: 0
Explanation: The string is already alternating.

Example 3:
Input: s = "1110"
Output: 1
Explanation: Use the second operation on the second element to make s = "1010".

Constraints:

1 <= s.length <= 105
s[i] is either '0' or '1'.
*/

// TimeLimitExceed
func minFlipsTwoSlow(s string) int {
	n := len(s)
	s += s
	var s1, s2 string
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			s1 += "0"
			s2 += "1"
		} else {
			s1 += "1"
			s2 += "0"
		}
	}

	var ans1, ans2 int
	ans := 1 << 31
	for i := 0; i < len(s); i++ {
		if s1[i] != s[i] {
			ans1++
		}
		if s2[i] != s[i] {
			ans2++
		}

		if i >= n {
			if s1[i-n] != s[i-n] {
				ans1--
			}
			if s2[i-n] != s[i-n] {
				ans2--
			}
		}

		if i >= n-1 {
			if ans > ans1 {
				ans = ans1
			}
			if ans > ans2 {
				ans = ans2
			}
		}
	}

	return ans
}

func minFlips(s string) int {
	var ns0 int // op start from '0'
	var ns1 int // op start from '1'
	ans := 1 << 31

	for i := 0; i < len(s)*2; i++ {
		if i < len(s) {
			if (i%2 == 0 && s[i] == '0') || (i%2 == 1 && s[i] == '1') {
				ns1++
			} else {
				ns0++
			}

			continue
		}

		iter := i % len(s)
		if (iter%2 == 0 && s[iter] == '0') || (iter%2 == 1 && s[iter] == '1') {
			ns1--
		} else {
			ns0--
		}

		if (i%2 == 0 && s[iter] == '0') || (i%2 == 1 && s[iter] == '1') {
			ns1++
		} else {
			ns0++
		}

		ans = min(ns0, ans)
		ans = min(ns1, ans)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
