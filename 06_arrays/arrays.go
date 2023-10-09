package arrays

import "fmt"

func Arrays() {

	fmt.Println("***************** ARRAYS ***************")

	// string array
	var s [2]string

	s[0] = "Chander"
	s[1] = "Parkash"

	fmt.Printf("\nat index 0 value is %s & at index 1 value is %s\n", s[0], s[1])

	// int array
	var i = [5]int{1, 2, 3, 4, 5}

	for _, v := range i {
		fmt.Printf("%d ", v)
	}
}
