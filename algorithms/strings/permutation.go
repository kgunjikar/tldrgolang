package strings

func permuteInt(s *[]int, l, r int, retVal *[][]int) {
	if l == r {
		*retVal = append(*retVal, (*s)[l:r]...)
	} else {

	}
}

func permutation(s []int) [][]int {
	retVal := make([][]int, 0)
	//permuteInt(&s, 0, len(s)-1, &retVal)
	return retVal
}
