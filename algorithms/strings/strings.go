package strings

import "fmt"

func findIndex(s string, c rune, m int) int {

	for k, val := range s {
		if val == c && k > m {
			fmt.Printf("%d %d \t", k, m)
			return k
		}
	}
	return -1
}

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	prev := -1
	for _, val := range s {
		temp := findIndex(t, val, prev)
		if temp == -1 {
			return false
		}
		if temp > prev {
			prev = temp
		} else {
			return false
		}

	}
	return true
}

func longestPalindrome(s string) int {

	runeMap := make(map[rune]int)
	for _, val := range s {
		runeMap[val] += 1
	}
	maxodd := 0
	count := 0
	for _, runeCount := range runeMap {
		if runeCount%2 == 0 {
			count += runeCount
		} else {
			count = count + runeCount - 1
			if runeCount > maxodd {
				maxodd = runeCount
			}
		}
	}
	if maxodd != 0 {
		count = count + 1
	}

	fmt.Printf("%d\n", count)
	return count
}
