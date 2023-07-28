package leetcode

import (
	"fmt"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	// [9,7,8]
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	// [10]
	fmt.Println(partitionLabels("eccbbbbdec"))
}
