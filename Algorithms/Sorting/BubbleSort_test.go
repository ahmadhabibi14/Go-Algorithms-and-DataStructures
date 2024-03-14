package sorting

import (
	"fmt"
	"testing"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 2, 8, 12, 3}

	fmt.Println("Array before sorting:", arr)

	bubbleSort(arr)

	fmt.Println("Array after sorting:", arr)
}
