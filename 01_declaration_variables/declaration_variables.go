package declaration_variables

import "fmt"

func DeclarationAndVariables() {

	fmt.Println("Declarations and Variables in Golang - start")

	var a int
	fmt.Println(a) // 0 (default)

	var b int = 5
	fmt.Println(b) // 5

	var c = 3
	fmt.Println(c) // c

	a = 12121
	fmt.Println(a) // 12121

	d := 12
	fmt.Println(d) // 12

	var e, f = 12, 14
	fmt.Println(e) // 12
	fmt.Println(f) // 14

	var (
		name = "chander"
		age  = 24
	)

	fmt.Println(name) // chander
	fmt.Println(age)  // 24

	fmt.Println("Declarations and Variables in Golang - start")
}
