package algorithms

func generateParenthesis(n int) []string {
	return helper(n, 0, 0, "")
}

func helper(n, lCnt, rCnt int, tmp string) []string {
	if rCnt > lCnt || lCnt > n {
		return []string{}
	}
	if lCnt+rCnt == n*2 {
		return append([]string{}, tmp)
	}

	res := []string{}
	res = append(res, helper(n, lCnt+1, rCnt, tmp+"(")...)
	res = append(res, helper(n, lCnt, rCnt+1, tmp+")")...)

	return res
}
