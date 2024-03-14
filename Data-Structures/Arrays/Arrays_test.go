package arrays

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// Declaration and initialization of an array
	var numbers [5]int

	// Assigning values to array elements
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Accessing array elements
	fmt.Println("Element at index 2:", numbers[2])

	// Iterating over array elements
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element at index %d: %d\n", i, numbers[i])
	}

	// Declaration and initialization of an array with values
	grades := [4]float64{98.5, 78.2, 86.7, 91.0}

	// Length of an array
	fmt.Println("Length of grades array:", len(grades))
}
