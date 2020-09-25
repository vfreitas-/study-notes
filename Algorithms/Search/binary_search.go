package search

import (
	"math"
)

/**
 * Binary Search Algorithm
 * 
 * Binary search compares the target value to the middle element of the array. 
 * If they are not equal, the half in which the target cannot lie is eliminated 
 * and the search continues on the remaining half, again taking the middle element
 * to compare to the target value, and repeating this until the target value is found. 
 * If the search ends with the remaining half being empty, the target is not in the array.
 *
 */
func binarySearchNumber (arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		sum := float64(left + right)
		mid := int(math.Floor(sum / 2))

		if (arr[mid] == target) {
			return mid
		}

		if (target < arr[mid]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func binarySearchString (arr []string, target string) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		sum := float64(left + right)
		mid := int(math.Floor(sum / 2))

		if (arr[mid] == target) {
			return mid
		}

		if (target < arr[mid]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
