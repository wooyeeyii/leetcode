package leetcode

/**
907. Sum of Subarray Minimums

Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous) subarray of arr. Since the answer may be large, return the answer modulo 10^9 + 7.

Example 1:
Input: arr = [3,1,2,4]
Output: 17
Explanation:
Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
Sum is 17.

Example 2:
Input: arr = [11,81,94,43,3]
Output: 444

Constraints:
1 <= arr.length <= 3 * 104
1 <= arr[i] <= 3 * 10^4
*/

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}
func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
func (s *stack) empty() bool {
	return len(*s) == 0
}
func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

// solution: https://leetcode.com/problems/sum-of-subarray-minimums/solutions/178876/stack-solution-with-very-detailed-explanation-step-by-step/
func sumSubarrayMins(arr []int) int {
	l := len(arr)
	// PLE to denote Previous Less Element.
	// NLE to denote Next Less Element.
	ple, nle := stack{}, stack{}
	left, right := make([]int, l, l), make([]int, l, l)
	for i := 0; i < l; i++ {
		left[i] = i + 1
		right[i] = l - i
	}

	for i := 0; i < l; i++ {
		for !ple.empty() && arr[ple.top()] > arr[i] {
			ple.pop()
		}
		if ple.empty() {
			left[i] = i + 1
		} else {
			left[i] = i - ple.top()
		}
		ple.push(i)

		// for next less
		for !nle.empty() && arr[nle.top()] > arr[i] {
			t := nle.top()
			nle.pop()
			right[t] = i - t
		}
		nle.push(i)
	}

	mod := int(1e9 + 7)
	ans := 0
	for i := 0; i < l; i++ {
		ans = (ans + arr[i]*left[i]*right[i]) % mod
	}
	return ans
}

// Time Limit Exceeded
func sumSubarrayMins_Slow(arr []int) int {
	mod := 1000000007
	l := len(arr)
	sum := 0
	for i := 0; i < l; i++ {
		min := arr[i]
		sum = (sum + min) % mod
		for j := i + 1; j < l; j++ {
			if arr[j] < min {
				min = arr[j]
			}
			sum = (sum + min) % mod
		}
	}
	return sum
}
