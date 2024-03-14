package searching

import (
	"fmt"
	"testing"
)

// binarySearch searches for a target value in a sorted array.
// It returns the index of the target if found, or -1 if not found.
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func TestBinarySearch(t *testing.T) {
	// Example usage
	arr := []int{2, 4, 7, 10, 14, 19, 22, 30}
	target := 14

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}
