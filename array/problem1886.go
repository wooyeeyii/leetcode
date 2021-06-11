package array

/**
1886. Determine Whether Matrix Can Be Obtained By Rotation

Given two n x n binary matrices mat and target, return true if it is possible
to make mat equal to target by rotating mat in 90-degree increments, or false otherwise.

Example 1:
Input: mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
Output: true
Explanation: We can rotate mat 90 degrees clockwise to make mat equal target.

Example 2:
Input: mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
Output: false
Explanation: It is impossible to make mat equal to target by rotating mat.

Example 3:
Input: mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
Output: true
Explanation: We can rotate mat 90 degrees clockwise two times to make mat equal target.

Constraints:

n == mat.length == target.length
n == mat[i].length == target[i].length
1 <= n <= 10
mat[i][j] and target[i][j] are either 0 or 1.
*/

func findRotation(mat [][]int, target [][]int) bool {
	if rotate0(mat, target) || rotate1(mat, target) || rotate2(mat, target) || rotate3(mat, target) {
		return true
	}
	return false

}

func rotate0(mat [][]int, target [][]int) bool {
	l := len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}

func rotate1(mat [][]int, target [][]int) bool {
	l := len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if mat[i][j] != target[j][l-1-i] {
				return false
			}
		}
	}
	return true
}

func rotate2(mat [][]int, target [][]int) bool {
	l := len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if mat[i][j] != target[l-1-i][l-1-j] {
				return false
			}
		}
	}
	return true
}

func rotate3(mat [][]int, target [][]int) bool {
	l := len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if mat[i][j] != target[l-1-j][i] {
				return false
			}
		}
	}
	return true
}
