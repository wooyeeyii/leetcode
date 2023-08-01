package leetcode

/**
128. Longest Consecutive Sequence

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Constraints:
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	longest := 0
	for k := range m {
		if !m[k-1] {
			l := 0
			for {
				if m[k+l] {
					l++
					continue
				}

				if l > longest {
					longest = l
				}
				break
			}
		}
	}
	return longest
}

// Time Limit Exceeded
func longestConsecutiveSlow(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 0
	min := 1000000000
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	longest := 0
	start, end := min, min
	for i := min; i <= max; i++ {
		if m[i] {
			end = i
		} else {
			if end-start+1 > longest {
				longest = end - start + 1
			}

			start = i + 1
		}
	}

	if end-start+1 > longest {
		longest = end - start + 1
	}

	return longest
}
