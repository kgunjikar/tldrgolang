package Whatistheoutput

import (
	"fmt"
)

/*
https://go.dev/doc/faq#nil_error
*/

type MyInterface interface {
	GetName() string
}
type MyInterfaceImpl struct {
	Name string
}

func (i *MyInterfaceImpl) GetName() string {
	return i.Name
}
func getNilMyInterface() MyInterface {
	return nil
}
func getNilMyInterfaceImpl() *MyInterfaceImpl {
	return nil
}

func test2() {
	v := getNilMyInterface()
	//fmt.Println(reflect.TypeOf(v).Name())
	if v == nil {
		if v = getNilMyInterfaceImpl(); v == nil {
			fmt.Println("Ok. This code is correct!")
		} else {
			fmt.Println("Opsss... It's not nil")
			//fmt.Println(v.GetName())
		}
	}
}
