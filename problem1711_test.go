package leetcode

import (
	"fmt"
	"testing"
)

func Test_countPairs(t *testing.T) {
	// 4
	fmt.Println(countPairs([]int{1, 3, 5, 7, 9}))
	// 15
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7}))
}
