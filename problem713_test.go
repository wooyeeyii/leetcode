package leetcode

import (
	"fmt"
	"testing"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	// 1
	fmt.Println(numSubarrayProductLessThanK([]int{57, 44, 92, 28, 66, 60, 37, 33, 52, 38, 29, 76, 8, 75, 22}, 18))
	// 0
	fmt.Println(numSubarrayProductLessThanK([]int{1, 1, 1}, 1))
	// 8
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	// 0
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
	// 6
	fmt.Println(numSubarrayProductLessThanK([]int{1, 1, 1}, 2))
}
