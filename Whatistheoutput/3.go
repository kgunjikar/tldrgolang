package Whatistheoutput

import "fmt"

// https://twitter.com/indradhanush92/status/1614230074647252993
func reuseTemp() {
	x := []int{1, 2, 3, 4}

	for _, item := range x {
		fmt.Println(&item)
	}
}
