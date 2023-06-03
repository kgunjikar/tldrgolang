package algorithms

var KEYS = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func combination(prefix string, digits string, idx int, res *[]string) {
	if idx >= len(digits) {
		*res = append(*res, prefix)
		return
	}
	index := digits[idx] - '0'
	letters := KEYS[index]
	for i := 0; i < len(letters); i++ {
		combination(prefix+string(letters[i]), digits, idx+1, res)
	}
}

func letterCombinations(digits string) []string {
	var res []string
	combination("", digits, 0, &res)
	return res
}
