package id33

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			// [low, mid] are orderd, [mid + 1, high] are not ordered
			if target >= nums[low] && target <= nums[mid] {
				if result := binarySearch(nums[low:mid+1], target); result != -1 {
					return result + low
				} else {
					return -1
				}
			} else {
				low = mid + 1
			}
		} else {
			// [mid, high] are ordered, [low, mid] are not ordered
			if target >= nums[mid] && target <= nums[high] {
				if result := binarySearch(nums[mid:high+1], target); result != -1 {
					return result + low
				} else {
					return -1
				}
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// binarySearch would looking for target in nums, return it's index if found, or -1 if not found
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// @lc code=end
