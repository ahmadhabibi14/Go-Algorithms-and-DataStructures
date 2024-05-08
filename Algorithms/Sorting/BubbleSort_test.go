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
	arr := []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}

	fmt.Println("Array before sorting:", arr)

	bubbleSort(arr)

	fmt.Println("Array after sorting:", arr)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
		bubbleSort(arr)
		fmt.Println("Sorted with BubbleSort \t:", arr)
	}
}