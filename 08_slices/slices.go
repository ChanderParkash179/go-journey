package slices

import "fmt"

func Slices() {
	fmt.Println("Slices : start")

	// slices comes with dynamic sizing

	// normal slice
	a := []int{1, 2, 3, 4}
	fmt.Println("slice a's values are: ", a)

	// using make
	b := make([]int, 4, 5)
	fmt.Println("values of slice (b): ", b)
	fmt.Println("length of slice (b): ", len(b))
	fmt.Println("capacity of slice (b): ", cap(b))

	b = append(b, 55, 66)
	fmt.Println("after appending 55 & 66 in slice (b): ", b)
	fmt.Println("Slices : end")
}
