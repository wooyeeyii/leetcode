package leetcode

import (
	"fmt"
	"testing"
)

func Test_strWithout3a3b(t *testing.T) {
	// "abb", "bab" and "bba" are all correct answers
	fmt.Println(strWithout3a3b(1, 2))
	// "aabaa"
	fmt.Println(strWithout3a3b(4, 1))
	// "bbabbabba"
	fmt.Println(strWithout3a3b(3, 6))
	// "bbabbab"
	fmt.Println(strWithout3a3b(2, 5))
	// "aabbaabb"
	fmt.Println(strWithout3a3b(4, 4))
}
