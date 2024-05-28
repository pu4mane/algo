package binarysearch

func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, len(nums)-1

	for l <= r {
		mid := ((r - l) / 2) + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return 0
}
