package array

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	nums := []int {1,2,3,4,5,6,7}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{-1,-100,3,99}
	rotate(nums, 2)
	fmt.Println(nums)
}
