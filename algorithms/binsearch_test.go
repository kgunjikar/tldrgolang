package algorithms

import "testing"

func Test_search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	if search(nums, 9) == -1 {
		t.Fail()
	}
}

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	if searchInsert(nums, 8) != 3 {
		t.Fail()
	}
}
