package binary

import (
	"fmt"
	"testing"
)

func Test_problem349(t *testing.T) {
	// 2
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	// 10
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func Test_problem349Example(t *testing.T) {
	// 2
	fmt.Println(singleNonDuplicateExample([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	// 10
	fmt.Println(singleNonDuplicateExample([]int{3, 3, 7, 7, 10, 11, 11}))
}
