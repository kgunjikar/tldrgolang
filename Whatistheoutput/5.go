package Whatistheoutput

import "fmt"

func Foo(myMap map[string]string, arr []int, a [4]int) {
	myMap["Hello"] = "Check"
	myMap["Best"] = "Check"
	arr[0] = 1
	arr = append(arr, 10)
	arr = append(arr, 11)

	a[2] = 6

}

func byValueOrReference() {
	myMap := make(map[string]string)
	myArr := make([]int, 2)
	intArr := [4]int{0, 1, 2, 4}
	fmt.Printf("Before:\n%v\n%v\n%v\n", myMap, myArr, intArr)
	Foo(myMap, myArr, intArr)
	fmt.Printf("After: \n%v\n%v\n%v\n", myMap, myArr, intArr)
}
