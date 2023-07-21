package leetcode

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func Test_removeNodes(t *testing.T) {
	n1 := &common.ListNode{Val: 5}
	n2 := &common.ListNode{Val: 2}
	n3 := &common.ListNode{Val: 13}
	n4 := &common.ListNode{Val: 3}
	n5 := &common.ListNode{Val: 8}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	r := removeNodes(n1)
	fmt.Println(r.ToString())
}

func Test_removeNodes_2(t *testing.T) {
	n1 := &common.ListNode{Val: 1}
	n2 := &common.ListNode{Val: 1}
	n3 := &common.ListNode{Val: 1}
	n4 := &common.ListNode{Val: 1}
	n5 := &common.ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	r := removeNodes(n1)
	fmt.Println(r.ToString())
}
