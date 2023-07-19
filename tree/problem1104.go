package tree

/*
1104. Path In Zigzag Labelled Binary Tree

In an infinite binary tree where every node has two children, the nodes are labelled in row order.
In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right,
while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.

Given the label of a node in this tree, return the labels in the path from the root of the tree to the node with that label.

Example 1:
Input: label = 14
Output: [1,3,4,14]

Example 2:
Input: label = 26
Output: [1,2,6,10,26]

Constraints:
1 <= label <= 10^6
*/

func pathInZigZagTree(label int) []int {
	if label == 0 {
		return []int{}
	}

	r := make([]int, 0, 0)
	end := 2
	for label > end-1 {
		r = append(r, end)
		end *= 2
	}

	r = append(r, end)
	pos := label
	l := len(r)
	if l%2 == 0 {
		pos = r[l-1] - 1 + r[l-2] - pos
	}

	for i := len(r) - 1; i > 0; i-- {
		cur := pos
		if i%2 == 1 {
			cur = r[i] - 1 + r[i-1] - cur
		}

		r[i] = cur
		pos = pos / 2
	}
	r[0] = 1

	return r
}
