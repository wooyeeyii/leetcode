package leetcode

/**
984. String Without AAA or BBB

Given two integers a and b, return any string s such that:
s has length a + b and contains exactly a 'a' letters, and exactly b 'b' letters,
The substring 'aaa' does not occur in s, and
The substring 'bbb' does not occur in s.

Example 1:
Input: a = 1, b = 2
Output: "abb"
Explanation: "abb", "bab" and "bba" are all correct answers.

Example 2:
Input: a = 4, b = 1
Output: "aabaa"

Constraints:
0 <= a, b <= 100
It is guaranteed such an s exists for the given a and b.
*/

func strWithout3a3b(a int, b int) string {
	ca, cb := byte('a'), byte('b')
	if b > a {
		ca, cb = 'b', 'a'
		a, b = b, a
	}

	l := a + b
	s := make([]byte, 0, 0)
	aC, bC := 0, 0
	aT, bT := 0, 0
	for i := 0; i < l; i++ {
		if aT < a && aC < 2 {
			s = append(s, ca)
			aC++
			aT++
		} else if bT < b && bC < 1 {
			s = append(s, cb)
			bC++
			bT++
		} else {
			if aT == a {
				break
			}

			aC, bC = 0, 0
			i -= 1
		}
	}

	r := make([]byte, l, l)
	j := b - bT
	k := 0
	for i := 0; i < len(s); i++ {
		r[k] = s[i]
		k++

		if s[i] == cb && j > 0 {
			r[k] = cb
			k++
			j--
		}
	}

	return string(r)
}
