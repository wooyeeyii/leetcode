package leetcode

func balancedStringSplit(s string) int {
	cntL, cntR, cnt := 0, 0, 0
	for _, c := range s {
		if c == 'R' {
			cntR++
		} else {
			cntL++
		}

		if cntR == cntL {
			cnt += 1
			cntR, cntL = 0, 0
		}
	}
	return cnt
}
