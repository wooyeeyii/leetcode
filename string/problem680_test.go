package string

import (
	"fmt"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
}
