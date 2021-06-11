package string

/**
459. Repeated Substring Pattern

Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

Example 1:
Input: s = "abab"
Output: true
Explanation: It is the substring "ab" twice.

Example 2:
Input: s = "aba"
Output: false

Example 3:
Input: s = "abcabcabcabc"
Output: true
Explanation: It is the substring "abc" four times or the substring "abcabc" twice.

Constraints:
1 <= s.length <= 104
s consists of lowercase English letters.

*/

import "strings"

func repeatedSubstringPattern(s string) bool {
	last := s[len(s)-1:]
	endIdx := strings.Index(s, last)
	for endIdx < len(s)/2 {
		sub := s[0 : endIdx+1]
		if repeated(s, sub) {
			return true
		}

		if nextIdx := strings.Index(s[endIdx+1:], last); nextIdx == -1 {
			break
		} else {
			endIdx = endIdx + nextIdx + 1
		}
	}
	return false
}

func repeated(s, sub string) bool {
	ns := len(s)
	nsub := len(sub)
	if ns%nsub != 0 || ns == nsub {
		return false
	}

	idx := 0
	for idx < ns {
		if sub != s[idx:idx+nsub] {
			return false
		}
		idx += nsub
	}
	return true
}

// simple solution
func repeatedSubstringPatternSimple(s string) bool {
	l := len(s)

	for i := 1; i < len(s); i++ {
		if l%i == 0 {
			repeatStr := strings.Repeat(s[:i], l/i)
			if repeatStr == s {
				return true
			}
		}
	}
	return false
}
