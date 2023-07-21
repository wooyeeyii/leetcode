package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToString() string {
	n := l
	str := ""
	for n != nil {
		str += fmt.Sprintf("%d ", n.Val)
		n = n.Next
	}

	return str
}
