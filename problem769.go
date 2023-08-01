package leetcode

/**
769. Max Chunks To Make Sorted

You are given an integer array arr of length n that represents a permutation of the integers in the range [0, n - 1].
We split arr into some number of chunks (i.e., partitions), and individually sort each chunk. After concatenating them, the result should equal the sorted array.
Return the largest number of chunks we can make to sort the array.

Example 1:
Input: arr = [4,3,2,1,0]
Output: 1
Explanation:
Splitting into two or more chunks will not return the required result.
For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1, 2], which isn't sorted.

Example 2:
Input: arr = [1,0,2,3,4]
Output: 4
Explanation:
We can split into two chunks, such as [1, 0], [2, 3, 4].
However, splitting into [1, 0], [2], [3], [4] is the highest number of chunks possible.

Constraints:
n == arr.length
1 <= n <= 10
0 <= arr[i] < n
All the elements of arr are unique.
*/

func maxChunksToSorted(arr []int) int {
	chunks := 0
	for l := 0; l < len(arr); l++ {
		r := arr[l]
		for ; l <= r; l++ {
			if arr[l] > r {
				r = arr[l]
			}
		}
		l--
		chunks++
	}
	return chunks
}

func maxChunksToSorted1(arr []int) int {
	cnt := 0
	max := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

		if max == i {
			cnt++
			if i+1 < len(arr) {
				max = arr[i+1]
			} else {
				max = -1
			}
		}
	}
	if max != -1 {
		cnt++
	}
	return cnt
}

func maxChunksToSortedExample(arr []int) int {
	chunks := 0
	for left := 0; left < len(arr); left++ {
		right := arr[left]
		for ; left <= right; left++ {
			if arr[left] > right {
				right = arr[left]
			}
		}
		left--
		chunks++
	}
	return chunks
}
