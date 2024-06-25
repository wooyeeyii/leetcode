package leetcode

/**
813. Largest Sum of Averages

You are given an integer array nums and an integer k. You can partition the array into at most k non-empty adjacent subarrays.
The score of a partition is the sum of the averages of each subarray.
Note that the partition must use every integer in nums, and that the score is not necessarily an integer.

Return the maximum score you can achieve of all the possible partitions. Answers within 10^-6 of the actual answer will be accepted.

Example 1:
Input: nums = [9,1,2,3,9], k = 3
Output: 20.00000
Explanation:
The best choice is to partition nums into [9], [1, 2, 3], [9]. The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
We could have also partitioned nums into [9, 1], [2], [3, 9], for example.
That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.

Example 2:
Input: nums = [1,2,3,4,5,6,7], k = 4
Output: 20.50000

Constraints:
1 <= nums.length <= 100
1 <= nums[i] <= 10^4
1 <= k <= nums.length
*/

func largestSumOfAverages(nums []int, k int) float64 {
	l := len(nums)
	cur := 0.0
	memo := make([][]float64, l+1)
	memo[0] = make([]float64, l+1)
	for i := 0; i < l; i++ {
		cur += float64(nums[i])

		memo[i+1] = make([]float64, l+1)
		memo[i+1][1] = cur / float64(i+1)
	}

	return search(l, k, nums, memo)
}

func search(l int, k int, nums []int, memo [][]float64) float64 {
	if memo[l][k] > 0 {
		return memo[l][k]
	}

	if l < k {
		return 0
	}

	cur := 0.0
	for i := l - 1; i > 0; i-- {
		cur += float64(nums[i])
		memo[l][k] = max(memo[l][k], search(i, k-1, nums, memo)+cur/float64(l-i))
	}
	return memo[l][k]
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
