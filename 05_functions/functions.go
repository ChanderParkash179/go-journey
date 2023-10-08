package functions

import (
	"fmt"
	"time"
)

func Functions() {

	fmt.Println("*******************************************")
	// simple function : void
	simple()

	fmt.Println("*******************************************")
	// parameterized : void
	parameterized("parameterized")

	fmt.Println("*******************************************")
	// return type : return
	fmt.Println(returnValue())

	fmt.Println("*******************************************")
	// named return value : return
	fmt.Println(namedReturn())

	fmt.Println("*******************************************")
	fmt.Println(multiReturn())
	// return multiple values : return

	fmt.Println("*******************************************")
	// passing address and returning : pointer return  (pending)

	fmt.Println("*******************************************")
	// anonymous function with return :  : type 1st
	var (
		add = func(a int, b int) int {
			return a + b
		}
	)

	a := 55
	b := 99
	fmt.Printf("\nAddition of %d & %d is %d", a, b, add(a, b))

	// anonymous function with return : type 2nd
	func(a int, b int) {
		fmt.Println("I am second type of anonymous function")
		fmt.Printf("\n value of a is %d & b is %d", a, b)
	}(55, 66)

	// anonymous function with return : type 3rd
	fmt.Println("I am third type of anonymous function",
		func(text string) string {
			return text
		}("I am from third anonymous"))

	// anonymous function with return : type 4th
	today := time.DateTime
	func() {
		fmt.Printf("\nToday %s", today)
	}()

	fmt.Println("\n*******************************************")
	// higher order user defined typed function
	v1 := higherOrderDefinedFunctionType(5)
	fmt.Println(v1(10))
}

// simple function : void
func simple() {
	fmt.Println("I am simple function")
}

// parameterized function : void
func parameterized(text string) {
	fmt.Printf("I am %s function\n", text)
}

// return value function : return type
func returnValue() int {
	fmt.Println("I am return type based function")
	return 10
}

// named return value : return
func namedReturn() (text string) {
	fmt.Println("I am named return value")
	text = "this is named return `text` variable"
	return
}

// return multiple values : return
func multiReturn() (age int, height float32) {
	age = 24
	height = 5.4
	fmt.Printf("\n my age is %d and my height is %f", age, height)
	return
}

// passing address and returning : pointer return  (pending)

type First func(x int) int

func higherOrderDefinedFunctionType(x int) First {
	return func(y int) int {
		return x + y
	}
}
