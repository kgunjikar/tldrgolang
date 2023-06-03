package snowflake

import "math"

func lowestBitPair(n, k int) (int, int) {

	if n < 0 {
		return math.MinInt32, math.MinInt32
	}
	min := math.MaxInt32
	expVal := 1 << n
	for x := math.MinInt32; x < math.MaxInt32; x++ {
		if x&(x+k)&expVal == expVal {
			if x < min {
				min = x
			}
		}
	}
	if min == math.MaxInt32 {
		return -1, -1
	}
	return min, min + k
}
