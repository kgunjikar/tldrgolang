package snowflake

import (
	"fmt"
	"testing"
)

func Test_largestSubGrid(t *testing.T) {
	nums := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	fmt.Println(largestSubGrid(nums, 4))
}
