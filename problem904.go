package leetcode

/**
904. Fruit Into Baskets

You are visiting a farm that has a single row of fruit trees arranged from left to right.
The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.
You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
Given the integer array fruits, return the maximum number of fruits you can pick.

Example 1:
Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.

Example 2:
Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2].
If we had started at the first tree, we would only pick from trees [0,1].

Example 3:
Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2].
If we had started at the first tree, we would only pick from trees [1,2].

Constraints:
1 <= fruits.length <= 10^5
0 <= fruits[i] < fruits.length
*/

func totalFruit(fruits []int) int {
	l := len(fruits)
	if l <= 2 {
		return l
	}

	max, total := 0, 0
	a := fruits[0]
	total += 1
	i := 1
	for ; i < l && fruits[i] == a; i++ {
		total++
	}
	if i == l {
		return total
	}

	b := fruits[i]
	lastN, lastC := b, 1
	total++
	for j := i + 1; j < l; j++ {
		if fruits[j] != a && fruits[j] != b {
			if total > max {
				max = total
			}
			total = lastC + 1
			a, b = lastN, fruits[j]
			lastN, lastC = fruits[j], 1
		} else {
			if fruits[j] == lastN {
				lastC++
			} else {
				lastN, lastC = fruits[j], 1
			}
			total++
		}

	}

	if total > max {
		max = total
	}
	return max
}
