package strings

import "strconv"

func RepeatingDigits(start, end int) []int {
	retVal := make([]int, 0)
	for i := start; i <= end; i++ {
		strI := strconv.Itoa(i)
		countMap := make(map[rune]int)
		for _, v := range strI {
			_, ok := countMap[v]
			if ok {
				retVal = append(retVal, i)
			} else {
				countMap[v]++
			}
		}
	}
	return retVal
}
