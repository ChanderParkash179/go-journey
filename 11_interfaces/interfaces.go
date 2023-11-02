package interfaces

import "fmt"

func Interfaces() {
	a := ali{"ali", "ali@email.com", 22}

	displayDetails(a)
}

type display interface {
	getName() string
	getEmail() string
	getAge() int
}

type ali struct {
	name  string
	email string
	age   int
}

func (a ali) getName() string {
	return a.name
}

func (a ali) getEmail() string {
	return a.email
}

func (a ali) getAge() int {
	return a.age
}

func displayDetails(d display) {
	fmt.Println(d)
	fmt.Println(d.getName())
	fmt.Println(d.getEmail())
	fmt.Println(d.getAge())
}
