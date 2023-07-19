package tree

import (
	"fmt"
	"testing"
)

func Test_pathInZigZagTree_1(t *testing.T) {
	r := pathInZigZagTree(16)
	// [1,3,4,15,16]
	fmt.Println(r)

	r = pathInZigZagTree(14)
	// [1,3,4,14]
	fmt.Println(r)

	r = pathInZigZagTree(26)
	// [1,2,6,10,26]
	fmt.Println(r)

	r = pathInZigZagTree(0)
	fmt.Println(r)

	r = pathInZigZagTree(1)
	fmt.Println(r)
}
