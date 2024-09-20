package sorting

import (
	"fmt"
	"testing"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//Swapping largest value and replacing its position with the smallest value
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}

	fmt.Println("Array before sorting:", arr)

	selectionSort(arr)

	fmt.Println("Array after sorting:", arr)
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
		selectionSort(arr)
		fmt.Println("Sorted with SelectionSort \t:", arr)
	}
}
