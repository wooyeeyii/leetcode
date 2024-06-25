package leetcode

import "sort"

/**
945. Minimum Increment to Make Array Unique

You are given an integer array nums. In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.
Return the minimum number of moves to make every value in nums unique.
The test cases are generated so that the answer fits in a 32-bit integer.

Example 1:
Input: nums = [1,2,2]
Output: 1
Explanation: After 1 move, the array could be [1, 2, 3].

Example 2:
Input: nums = [3,2,1,2,1,7]
Output: 6
Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
It can be shown with 5 or less moves that it is impossible for the array to have all unique values.

Constraints:
1 <= nums.length <= 10^5
0 <= nums[i] <= 10^5
*/

func minIncrementForUnique(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int)
	for _, n := range nums {
		m[n] = m[n] + 1
	}

	cnts := make([][2]int, 0, 0)
	for k, v := range m {
		cnts = append(cnts, [2]int{k, v})
	}
	sort.SliceStable(cnts, func(i, j int) bool {
		return cnts[i][0] < cnts[j][0]
	})

	moves := 0
	n1 := cnts[0][0]
	c1 := cnts[0][1] - 1
	for i := 1; i < len(cnts); i++ {
		n2 := cnts[i][0]
		c2 := cnts[i][1] - 1
		if n2-n1-1 >= c1 {
			moves += (1 + c1) * c1 / 2
		} else {
			l := n2 - n1
			moves += l * (l - 1) / 2
			moves += (c1 - l + 1) * l
			c2 += c1 - (n2 - n1 - 1)
		}

		n1 = n2
		c1 = c2
	}

	moves += (1 + c1) * c1 / 2
	return moves
}

func minIncrementForUniqueExample(A []int) int {
	if len(A) < 2 {
		return 0
	}
	res := 0
	sort.Ints(A)

	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			res += A[i-1] + 1 - A[i] //the incremental value is just the number of moves
			A[i] = A[i-1] + 1
		}
	}
	return res
}
