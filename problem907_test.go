package leetcode

import (
	"fmt"
	"testing"
)

func Test_sumSubarrayMins(t *testing.T) {
	// 17
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	// 444
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}
