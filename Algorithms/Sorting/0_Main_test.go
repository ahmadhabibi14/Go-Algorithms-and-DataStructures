package sorting

import (
	"fmt"
	"testing"
)

func ToArray(array []string) {
	if len(array) >= 2 {
		array[0], array[1] = array[1], array[0]
	}
}

func TestYeay(t *testing.T) {
	var names []string = []string{"Habi", "Doni", "Dian"}
	var nPoints *[]string = &names
	
	fmt.Println("Before\t\t:", names)
	ToArray(names)

	fmt.Println("After\t\t:", names)
	fmt.Println("Pointered\t:", *nPoints)

	var ha = 5 / 2
	fmt.Println("HA", ha)
}