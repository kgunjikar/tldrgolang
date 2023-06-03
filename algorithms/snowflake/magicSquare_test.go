package snowflake

import "testing"

func Test_formingMagicSquare(t *testing.T) {
	s := []int{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}
	formingMagicSquare(s)
}
