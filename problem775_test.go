package leetcode

import (
	"fmt"
	"testing"
)

func Test_isIdealPermutation(t *testing.T) {
	// false
	fmt.Println(isIdealPermutation([]int{2, 1, 0}))
	// true
	fmt.Println(isIdealPermutation([]int{1, 0, 2}))
	// false
	fmt.Println(isIdealPermutation([]int{1, 2, 0}))
}
