package array

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	test1 := []int{1, 2, 3, 0, 0, 0}
	merge(test1, 3, []int{2, 5, 6}, 3)
	fmt.Println(test1)

	test2 := []int{1}
	merge(test2, 1, []int{}, 0)
	fmt.Println(test2)

	test3 := []int{0}
	merge(test3, 0, []int{1}, 1)
	fmt.Println(test3)
}
