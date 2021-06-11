package array

import (
	"fmt"
	"testing"
)

func Test_minFlips(t *testing.T) {
	// 2
	fmt.Println(minFlips("111000"))
	// 0
	fmt.Println(minFlips("101"))
	// 1
	fmt.Println(minFlips("1110"))
}
