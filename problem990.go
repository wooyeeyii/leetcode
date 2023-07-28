package leetcode

/*
*
990. Satisfiability of Equality Equations

You are given an array of strings equations that represent relationships between variables where
each string equations[i] is of length 4 and takes one of two different forms: "xi==yi" or "xi!=yi".
Here, xi and yi are lowercase letters (not necessarily different) that represent one-letter variable names.
Return true if it is possible to assign integers to variable names so as to satisfy all the given equations, or false otherwise.

Example 1:
Input: equations = ["a==b","b!=a"]
Output: false
Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.
There is no way to assign the variables to satisfy both equations.

Example 2:
Input: equations = ["b==a","a==b"]
Output: true
Explanation: We could assign a = 1 and b = 1 to satisfy both equations.

Constraints:
1 <= equations.length <= 500
equations[i].length == 4
equations[i][0] is a lowercase letter.
equations[i][1] is either '=' or '!'.
equations[i][2] is '='.
equations[i][3] is a lowercase letter.
*/

type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(num int) *UnionFind {
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = i
	}
	return &UnionFind{
		arr,
		num,
	}
}

func (u *UnionFind) Union(a, b int) bool {
	parentA := u.Find(a)
	parentB := u.Find(b)
	if parentA != parentB {
		u.parent[parentA] = parentB
		u.count--
		return true
	}
	return false
}

func (u *UnionFind) Find(a int) int {
	if u.parent[a] == a {
		return a
	}
	u.parent[a] = u.Find(u.parent[a])
	return u.parent[a]
}

func equationsPossibleExample(equations []string) bool {
	uf := NewUnionFind(26)
	for _, eq := range equations {
		if eq[1] == '=' {
			uf.Union(int(eq[0]-'a'), int(eq[3]-'a'))
		}
	}
	for _, eq := range equations {
		if eq[1] == '!' {
			if uf.Find(int(eq[0]-'a')) == uf.Find(int(eq[3]-'a')) {
				return false
			}
		}
	}
	return true
}

func equationsPossible(equations []string) bool {
	notEqual := []string{}
	var arr [26]uint8
	for i := 0; i < 26; i++ {
		arr[i] = uint8(i)
	}

	rootParent := func(p uint8) uint8 {
		for arr[p] != p {
			p = arr[p]
		}
		return p
	}

	for _, e := range equations {
		if e[1] == '!' {
			notEqual = append(notEqual, e)
		} else {
			a := rootParent(e[0] - 'a')
			b := rootParent(e[3] - 'a')
			arr[a] = b
		}
	}

	for _, e := range notEqual {
		a := rootParent(e[0] - 'a')
		b := rootParent(e[3] - 'a')
		if a == b {
			return false
		}
	}

	return true
}
