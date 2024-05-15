package sorting

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	printQuickSort()
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr []int = []int{12, 11, 13, 5, 6, 4, 7, 10, 2, 3, 1, 9, 8}
		arr = QuickSort(arr)
		fmt.Println("Array after sort:\t", arr)
	}
}

func printQuickSort() {
	var arr []int = []int{12, 11, 13, 5, 6, 4, 7, 10, 2, 3, 1, 9, 8}
	fmt.Println("Array before sort:\t", arr)

	arr = QuickSort(arr)
	fmt.Println("Array after sort:\t", arr)
}

func QuickSort(arr []int) []int {
	newArr := make([]int, len(arr))

	copy(newArr, arr)

	sort(newArr, 0, len(arr)-1)

	return newArr
}

func sort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}