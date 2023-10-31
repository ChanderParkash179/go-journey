package typecasting

import (
	"fmt"
	"strconv"
)

func TypeCasting() {
	fmt.Println("type casting : start")

	fmt.Println("*********************")

	// string to int
	var s string = "22"
	v, _ := strconv.Atoi(s)
	fmt.Printf("\nI am after converting from string to integer: %d", v)

	// int to float
	var f float32 = 34.5
	fi := int(f)
	fmt.Printf("\nI am after converting from float to int: %d", fi)

	// string to bytes
	str := s
	var b []byte = []byte(str)
	fmt.Printf("\nI am after converting from string to byte: %b", b)

	// byte to string
	var ss = string(b)
	fmt.Printf("\nI am after converting from byte to string: %s", ss)

	fmt.Println(b)

	fmt.Println("\n*********************")

	fmt.Println("type casting : end")
}
