package pointers

import "fmt"

func Pointers() {
	num := 22
	ptr := &num

	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
