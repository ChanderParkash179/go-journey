package structs

import "fmt"

func Structs() {
	a := ali{"ali", "ali@email.com", 24}

	display(a)

}

type ali struct {
	name  string
	email string
	age   int
}

func display(a ali) {
	fmt.Println(a)
	fmt.Println(a.name)
	fmt.Println(a.email)
	fmt.Println(a.age)
}
