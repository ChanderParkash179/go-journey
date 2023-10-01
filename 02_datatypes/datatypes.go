package datatypes

import (
	"fmt"
)

func DataTypes() {

	// boolean
	var isValid bool = false
	fmt.Println(isValid)

	// int, int8, int16, int32, int64
	var a int = 4
	var b int8 = 8
	var c int16 = 16
	var d int32 = 32
	var e int64 = 64

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// uint, uint8, uint16, uint32, uint64
	var aa uint = 4
	var bb uint8 = 8
	var cc uint16 = 16
	var dd uint32 = 32
	var ee uint64 = 64

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)

	// float32, float64
	var aaa float32 = 32.32
	var bbb float64 = 64.64

	fmt.Println(aaa)
	fmt.Println(bbb)

	// string
	var f string = "name"

	fmt.Println(f)

	// byte
	var g byte = 'b'
	fmt.Println(g)

	// pointers
	var value int = 20
	var ptr *int = &value

	fmt.Printf("\ndirect value is: %d", value)
	fmt.Printf("\nfrom pointer value is: %d", *ptr)
	fmt.Printf("\npointer address is: %v", ptr)

	// rune
	// complex64/128
	// uintptr
}
