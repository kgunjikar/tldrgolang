package snowflake

import (
	"fmt"
	"math"
)

//https://leetcode.com/discuss/interview-question/850974

func largestSubGrid(nums [][]int, maxSum int) int {
	n := len(nums)
	m := len(nums[0])

	if n != m {
		return math.MaxInt32
	}
	nosofSubGrid := 0
	for gridSize := 1; gridSize <= n; gridSize++ {
		fmt.Printf("gridSize: %d\n", gridSize)
		for startRow := 0; startRow < n; startRow++ {
			for startCol := 0; startCol < n; startCol++ {
				sum := 0
				for i := startRow; i < startRow+gridSize; i++ {
					for j := startCol; j < startCol+gridSize; j++ {
						if i >= n || j >= n {
							continue
						}
						fmt.Println(i, j)
						sum += nums[i][j]
					}
				}
				if sum < maxSum {
					nosofSubGrid++
				}
			}
		}
	}
	return nosofSubGrid
}
