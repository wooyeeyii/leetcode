package leetcode

import (
	"leetcode/common"
)

/*
487. Remove Nodes From Linked List

You are given the head of a linked list.
Remove every node which has a node with a strictly greater value anywhere to the right side of it.
Return the head of the modified linked list.

Example 1:
Input: head = [5,2,13,3,8]
Output: [13,8]
Explanation: The nodes that should be removed are 5, 2 and 3.
- Node 13 is to the right of node 5.
- Node 13 is to the right of node 2.
- Node 8 is to the right of node 3.

Example 2:
Input: head = [1,1,1,1]
Output: [1,1,1,1]
Explanation: Every node has value 1, so no nodes are removed.

Constraints:
The number of the nodes in the given list is in the range [1, 10^5].
1 <= Node.val <= 10^5
*/

// Optimize: use stack

func removeNodes_1(head *common.ListNode) *common.ListNode {
	nodes := make([]*common.ListNode, 0, 0)
	n := head
	for n != nil {
		nodes = append(nodes, n)
		n = n.Next
	}

	i, j := len(nodes)-2, len(nodes)-1
	for i >= 0 {
		if nodes[i].Val < nodes[j].Val {
			nodes[i] = nil
		} else {
			j = i
		}
		i--
	}

	var start *common.ListNode
	startIdx := 0
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			start = nodes[i]
			startIdx = i
			nodes[i].Next = nil
			break
		}
	}

	i, j = startIdx, startIdx+1
	for j < len(nodes) {
		if nodes[j] != nil {
			nodes[j].Next = nil
			nodes[i].Next = nodes[j]
			i = j
		}
		j++
	}
	return start
}

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(c interface{}) {
	*s = append(*s, c) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func removeNodes(head *common.ListNode) *common.ListNode {
	n := head
	stack := new(Stack)
	for n != nil {
		top, v := stack.Pop()
		for v && top.(*common.ListNode).Val < n.Val {
			top, v = stack.Pop()
		}

		if v {
			stack.Push(top)
		}
		stack.Push(n)
		n = n.Next
	}

	var next *common.ListNode
	start, v := stack.Pop()
	for v {
		sn := start.(*common.ListNode)
		sn.Next = next
		next = sn
		start, v = stack.Pop()
	}
	return next
}
