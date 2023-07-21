package leetcode

import (
	"fmt"
	"testing"
)

func Test_totalFruit(t *testing.T) {
	// 3
	fmt.Println(totalFruit([]int{1, 2, 1}))
	// 3
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	// 4
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
}
