package binary

/*
540. Single Element in a Sorted Array

You are given a sorted array consisting of only integers where every element appears exactly twice,
except for one element which appears exactly once. Find this single element that appears only once.

Follow up: Your solution should run in O(log n) time and O(1) space.

Example 1:
Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2

Example 2:
Input: nums = [3,3,7,7,10,11,11]
Output: 10

Constraints:
1 <= nums.length <= 10^5
0 <= nums[i] <= 10^5

*/

// O(n) time and O(1) space
func singleNonDuplicate(nums []int) int {
	var res int
	for _, n := range nums {
		res = res ^ n
	}
	return res
}

func singleNonDuplicateExample(nums []int) int {
	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid%2 == 0 {
			// mid is even
			if nums[mid] == nums[mid+1] {
				lo = mid + 2
			} else {
				hi = mid
			}
		} else {
			// mid is odd
			if nums[mid] == nums[mid-1] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}
	return nums[lo]
}

func singleNonDuplicateExample2(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := lo + (hi-lo)/2
		n := mid ^ 1 // if even, mid + 1; if odd, mid - 1
		if nums[mid] == nums[n] {
			// if mid is even, then nums[mid] = nums[mid + 1], single number is >= mid + 2
			// if mid is odd, then nums[mid] = nums[mid - 1], single number is >= mid + 1
			// so we choose mid + 1
			lo = mid + 1
		} else {
			// maybe nums[hi] is the single numer or
			// maybe the single number is to the left of nums[hi]
			// <= hi
			hi = mid
		}
	}
	return nums[lo]
}
