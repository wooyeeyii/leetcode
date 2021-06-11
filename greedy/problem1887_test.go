package greedy

import (
	"fmt"
	"testing"
)

func Test_reductionOperations(t *testing.T) {
	// 0
	fmt.Println(reductionOperations([]int{1, 1, 1}))
	// 3
	fmt.Println(reductionOperations([]int{5, 1, 3}))
	// 4
	fmt.Println(reductionOperations([]int{1, 1, 2, 2, 3}))
}
