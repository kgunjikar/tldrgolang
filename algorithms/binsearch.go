package algorithms

func search(nums []int, target int) int {
	mid := len(nums) / 2
	var start, end int

	start = 0
	end = len(nums)
	for {
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			start = mid
		} else {
			end = mid
		}
		if start > end {
			return -1
		} else {
			mid = (start + end) / 2
		}

	}
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left >= len(nums) {
		left = left - 1
	}
	return left
}
