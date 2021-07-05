package string

/*
680. Valid Palindrome II

Given a string s, return true if the s can be palindrome after deleting at most one character from it.

Example 1:
Input: s = "aba"
Output: true

Example 2:
Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.

Example 3:
Input: s = "abc"
Output: false

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters.
*/

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for ; l <= r; {
		if s[l] != s[r] {
			return isPalindrome(s[l+1:r+1]) || isPalindrome(s[l:r])
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for ; l < r; {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
