package types

import (
	"reflect"
)

// int is same as the ARCH size
func intPromo() (string, string, string, string) {
	var test int = 0
	test1 := 1 * int32(test)
	test2 := 2 * int64(test)
	return reflect.TypeOf(1).String(), reflect.TypeOf(test).String(), reflect.TypeOf(test1).String(), reflect.TypeOf(test2).String()
}

func floatPromo() (string, string, string) {
	var test int = 0
	test1 := 0.1619 * float64(test)
	test2 := 2.0 * float32(test)
	return reflect.TypeOf(1.0).String(), reflect.TypeOf(test1).String(), reflect.TypeOf(test2).String()
}
