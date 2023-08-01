package leetcode

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	// 0
	fmt.Println(longestConsecutive([]int{}))
	// 5
	fmt.Println(longestConsecutive([]int{1, 3, 5, 2, 4}))
	// 4
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	// 9
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
