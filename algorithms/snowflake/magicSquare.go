package snowflake

import (
	"fmt"
	"math"
)

func swapAllRows(arr [][]int, allMagics *[][][]int) {
	*allMagics = append(*allMagics, arr)

	temp := make([]int, 3)
	copy(temp, arr[0])
	copy(arr[0], arr[1])
	copy(arr[1], temp)
	fmt.Println(arr)

	*allMagics = append(*allMagics, arr)

	copy(temp, arr[1])
	copy(arr[1], arr[2])
	copy(arr[2], temp)
	fmt.Println(arr)

	*allMagics = append(*allMagics, arr)

}

func swapAllCols(arr [][]int, allMagics *[][][]int) {

	for i := 0; i < 3; i++ {
		temp := arr[i][0]
		arr[1][0] = arr[i][1]
		arr[i][1] = temp
	}
	fmt.Println(arr)
	*allMagics = append(*allMagics, arr)

	for i := 0; i < 3; i++ {
		temp := arr[i][0]
		arr[1][0] = arr[i][2]
		arr[i][2] = temp
	}
	fmt.Println(arr)

	*allMagics = append(*allMagics, arr)

}

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func generateMagicMatrix(arr [][]int, allMagics *[][][]int) {

	swapAllRows(arr, allMagics)
	swapAllCols(arr, allMagics)

	temp := arr[0][0]
	arr[0][0] = arr[2][2]
	arr[2][2] = temp
	fmt.Println(arr)

	*allMagics = append(*allMagics, arr)

	temp = arr[0][2]
	arr[0][2] = arr[2][0]
	arr[2][0] = temp
	fmt.Println(arr)

	*allMagics = append(*allMagics, arr)
}

func formingMagicSquare(s [][]int) int {
	// Write your code here

	allMagics := make([][][]int, 0)
	oneMagic := [][]int{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}}
	generateMagicMatrix(oneMagic, &allMagics)

	minCost := math.MaxInt32
	for _, v := range allMagics {
		cost := diff(v, s)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost

}
