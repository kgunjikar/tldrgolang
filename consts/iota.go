package consts

import "fmt"

const (
	Write = iota
	Read  = iota
	Execute
)

func IotaCheck() {
	fmt.Println(Write, Read, Execute)
}
