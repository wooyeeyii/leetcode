package string

import (
	"fmt"
	"testing"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcdabcd"))
	fmt.Println(repeatedSubstringPattern("aaaaaaa"))
}
