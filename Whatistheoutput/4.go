package Whatistheoutput

import "fmt"

// https://github.com/golang/go/issues/57759
func copyChecks() {
	x := []int{1, 2, 3, 4}

	copy(x, x) // will avoid copy
	//	_ = append(x[:0], x) // will do the copy
}

func lenMap() {
	myMap := make(map[string]struct{})
	arr := make([]int, 0)
	fmt.Println(len(myMap), len(arr))
}
