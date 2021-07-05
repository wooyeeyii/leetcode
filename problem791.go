package leetcode

import "strings"

/*
791. Custom Sort String

order and str are strings composed of lowercase letters. In order, no letter occurs more than once.
order was sorted in some custom order previously. We want to permute the characters of str so that they
match the order that order was sorted. More specifically, if x occurs before y in order, then x should occur before y in the returned string.

Return any permutation of str (as a string) that satisfies this property.

Example:
Input:
order = "cba"
str = "abcd"
Output: "cbad"
Explanation:
"a", "b", "c" appear in order, so the order of "a", "b", "c" should be "c", "b", and "a".
Since "d" does not appear in order, it can be at any position in the returned string. "dcba", "cdba", "cbda" are also valid outputs.

Note:
order has length at most 26, and no character is repeated in order.
str has length at most 200.
order and str consist of lowercase letters only.
*/

func customSortString(order string, str string) string {
	var res string
	cnts := [26]int{}
	for i := 0; i < len(str); i++ {
		cnts[str[i]-'a']++
	}

	for i := 0; i < len(order); i++ {
		s := order[i]
		if cnts[s-'a'] > 0 {
			res += strings.Repeat(string(s), cnts[s-'a'])
		}
		cnts[s-'a'] = 0
	}

	for i := 0; i < 26; i++ {
		if cnts[i] > 0 {
			res += strings.Repeat(string(rune(i+'a')), cnts[i])
		}
	}
	return res
}
