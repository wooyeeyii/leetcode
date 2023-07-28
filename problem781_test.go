package leetcode

import (
	"fmt"
	"testing"
)

func Test_numRabbits(t *testing.T) {
	//6
	fmt.Println(numRabbits([]int{0, 0, 1, 1, 1}))
	// 5
	fmt.Println(numRabbits([]int{1, 0, 1, 0, 0}))
	// 5
	fmt.Println(numRabbits([]int{1, 1, 2}))
	// 11
	fmt.Println(numRabbits([]int{10, 10, 10}))
}
