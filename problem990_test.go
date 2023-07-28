package leetcode

import (
	"fmt"
	"testing"
)

func Test_equationsPossible(t *testing.T) {
	// false
	fmt.Println(equationsPossible([]string{"a==b", "e==c", "b==c", "a!=e"}))
	//false
	fmt.Println(equationsPossible([]string{"b!=f", "c!=e", "f==f", "d==f", "b==f", "a==f"}))
	// false
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
	// true
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))
}
