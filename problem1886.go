package leetcode

func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		if matrixEquals(mat, target) {
			return true
		}
		mat = rotation(mat)
	}
	return false
}

func rotation(mat [][]int) [][]int {
	n := len(mat)
	mat2 := make([][]int, n, n)
	for i := 0; i < n; i++ {
		mat2[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			mat2[i][j] = mat[j][n-1-i]
		}
	}
	return mat2
}

func matrixEquals(mat1 [][]int, mat2 [][]int) bool {
	n := len(mat1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat1[i][j]^mat2[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func findRotationExample(a [][]int, b [][]int) bool {
	n := len(a)
	c90, c180, c270, c0 := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[i][j] == a[n-j-1][i] {
				c90++
			}
			if b[i][j] == a[n-i-1][n-j-1] {
				c180++
			}
			if b[i][j] == a[j][n-i-1] {
				c270++
			}
			if b[i][j] == a[i][j] {
				c0++
			}
		}
	}

	if c90 == n*n || c270 == n*n || c180 == n*n || c0 == n*n {
		return true
	}

	return false
}
