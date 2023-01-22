package slices

import (
	"fmt"
	"sort"
)

func sliceImmutable() {

	var arr []int
	_ = append(arr, 3)
	fmt.Println(arr)
}

// https://www.reddit.com/r/golang/comments/10b4ofx/confused_about_array_and_slices/
/*
It seems you've already got an answer.
As Go does not hold too much on Principle of least surprise, slc = arr[:] assignment
expression does not copy delimited elements (unlike f.E. Python),
but rather creates a slice (a "window") to original array (pointer to source array values, length and capacity).

Expected result can be achieved by using copy builtin function.

example


*/
func slicesVsArrays() {
	var arr = [3]int{3, 1, 2}
	var slc = arr[:]

	sort.Ints(slc)
	fmt.Printf("val1: %v \nval2: %v \nptr1: %p \nptr2: %p", arr, slc, &arr, &slc)
}
