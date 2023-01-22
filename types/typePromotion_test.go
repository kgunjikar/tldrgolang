package types

import (
	"fmt"
	"testing"
)

func Test_floatPromo(t *testing.T) {
	a, b, c := floatPromo()
	fmt.Println(a, b, c)
}

func Test_intPromo(t *testing.T) {
	a, b, c, d := intPromo()
	fmt.Println(a, b, c, d)
}
