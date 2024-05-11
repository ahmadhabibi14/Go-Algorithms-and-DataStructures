package sorting

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	printShellSort()
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printShellSort()
	}
}

func shellSort(arr []int) {
	n := len(arr)

	// Start with a big gap, then reduce the gap
	for gap := n / 2; gap > 0; gap /= 2 {
		// Do a gapped insertion sort for this gap size.
		// The first gap elements arr[0..gap-1] are already in gapped order
		// keep adding one more element until the entire array is gap sorted
		for i := gap; i < n; i++ {
			// add arr[i] to the elements that have been gap sorted
			// save arr[i] in temp and make a hole at position i
			temp := arr[i]
			// shift earlier gap-sorted elements up until the correct location for arr[i] is found
			var j int
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}
			// put temp (the original arr[i]) in its correct location
			arr[j] = temp
		}
	}
}

func printShellSort() {
	var arr []int = []int{12, 11, 13, 5, 6, 4, 7, 10, 2, 3, 1, 9, 8}
	fmt.Println("Array before sort:\t", arr)

	shellSort(arr)
	fmt.Println("Array after sort:\t", arr)
}