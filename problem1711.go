package leetcode

/*
1711. Count Good Meals

A good meal is a meal that contains exactly two different food items with a sum of deliciousness equal to a power of two.
You can pick any two different foods to make a good meal.
Given an array of integers deliciousness where deliciousness[i] is the deliciousness of the deliciousness where deliciousness[i] item of food,
return the number of different good meals you can make from this list modulo 10^9 + 7.
Note that items with different indices are considered different even if they have the same deliciousness value.

Example 1:
Input: deliciousness = [1,3,5,7,9]
Output: 4
Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.

Example 2:
Input: deliciousness = [1,1,1,3,3,3,7]
Output: 15
Explanation: The good meals are (1,1) with 3 ways, (1,3) with 9 ways, and (1,7) with 3 ways.

Constraints:
1 <= deliciousness.length <= 10^5
0 <= deliciousness[i] <= 2^20
*/
func countPairs(deliciousness []int) int {
	mod := int64(1000000007)
	cnt := int64(0)
	m := map[int]int64{}
	for _, n := range deliciousness {
		pow := 1
		for i := 0; i < 22; i++ {
			if pow-n >= 0 {
				cnt += m[pow-n]
				cnt %= mod
			}
			pow *= 2
		}

		if c, ok := m[n]; ok {
			m[n] = c + 1
		} else {
			m[n] = 1
		}
	}
	return int(cnt)
}

// Time Limit Exceeded
func countPairs_TLE(deliciousness []int) int {
	cnt := 0
	l := len(deliciousness)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if isPowersOf2(deliciousness[i] + deliciousness[j]) {
				cnt++
			}
		}
	}
	return cnt
}

func isPowersOf2(n int) bool {
	if n == 0 {
		return false
	}

	for n > 1 {
		if n&0x1 == 1 {
			return false
		}
		n = n >> 1
	}
	return true
}
