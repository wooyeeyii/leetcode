package leetcode

import (
	"fmt"
	"testing"
)

func Test_minIncrementForUnique(t *testing.T) {
	// 16
	fmt.Println(minIncrementForUnique([]int{7, 2, 7, 2, 1, 4, 3, 1, 4, 8}))
	// 1
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	// 6
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
