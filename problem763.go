package leetcode

/**
763. Partition Labels

You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.
Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.
Return a list of integers representing the size of these parts.

Example 1:
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

Example 2:
Input: s = "eccbbbbdec"
Output: [10]

Constraints:
1 <= s.length <= 500
s consists of lowercase English letters.
*/

func partitionLabels(s string) []int {
	var cnt [][2]int
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; !ok {
			cnt = append(cnt, [2]int{i, i})
			m[s[i]] = len(cnt) - 1
		} else {
			cnt[v][1] = i
		}
	}

	ls := make([]int, 0, 0)
	if len(cnt) == 0 {
		return ls
	}

	min, max := cnt[0][0], cnt[0][1]
	for i := 1; i < len(cnt); i++ {
		p := cnt[i]
		if p[0] > max {
			ls = append(ls, max-min+1)
			min = p[0]
			max = p[1]
		} else {
			if p[1] > max {
				max = p[1]
			}
		}
	}
	ls = append(ls, max-min+1)
	return ls
}
