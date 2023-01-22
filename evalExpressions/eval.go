package evalExpressions

import "fmt"

func eval1() {
	var i, j, k, l int = 2, 3, 0, 0

	var a, b float64

	k = i / j * j
	l = j / i * i
	a = float64(i / j * j)
	b = float64(j / i * i)

	fmt.Println(k, l, a, b)

}

func eval2() {

	var a, b int

	a = -3 - -25
	b = -3 - -(-3)
	fmt.Println(a, b)
}

/*
func eval3() {
	var a, b float32 = 5, 3

	c := a % b
	fmt.Println(c)

}

func evalEqual() {
	var x,y int = 10, 20

	if x == y ;
	fmt.Println(x,y)
}

func evalEqual1() {
	var x int = 3
	var y float32 = 3.0
	if x == y {

		fmt.Println("Test", x, y)
	} else {
		fmt.Println(x, y)
	}

}

func evalEqual2() {
	var x, y, z int = 3, 0, 0

	y = x = 10
	z = x < 10
} */