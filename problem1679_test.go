package leetcode

import (
	"fmt"
	"testing"
)

func Test_maxOperations(t *testing.T) {
	// 2
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	// 1
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}
