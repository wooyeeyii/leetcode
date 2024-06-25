package leetcode

import (
	"fmt"
	"testing"
)

func Test_largestSumOfAverages(t *testing.T) {
	// 20.000
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
	// 20.50000
	fmt.Println(largestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4))
}
