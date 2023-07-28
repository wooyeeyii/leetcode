package leetcode

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	// "fl"
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	// ""
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
