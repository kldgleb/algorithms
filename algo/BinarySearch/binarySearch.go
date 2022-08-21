package binarysearch

// Есть упорядоченный массив целых чисел arr, нужно определить, есть ли в нём число X.
func BinarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := (right + left) / 2
		if arr[middle] == target {
			return true
		} else if arr[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}

	}
	return false
}

// Given a sorted array of distinct integers and a target value,
//  return the index if the target is found. If not,
//  return the index where it would be if it were inserted in order.

func SearchInsert(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	for left <= right {
		mid = (right + left) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	if nums[mid] > target {
		return mid
	}
	return mid + 1
}
