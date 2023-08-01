package leetcode

import (
	"fmt"
	"testing"
)

func Test_maxChunksToSorted(t *testing.T) {
	// 1
	fmt.Println(maxChunksToSorted([]int{1, 2, 3, 4, 5, 0}))
	// 4
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
	// 1
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	// 1
	fmt.Println(maxChunksToSorted([]int{4, 3, 1, 2, 0}))
	// 3
	fmt.Println(maxChunksToSorted([]int{1, 0, 3, 2, 4}))
}
