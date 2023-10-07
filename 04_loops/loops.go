package loops

import "fmt"

func Loops() {
	fmt.Println("I am for loop chapter")
	fmt.Println("**********************")

	// pure for loop
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("\npure for loop [sum is %d]", sum)
	fmt.Println("\n**********************")

	// while in shape of for
	fmt.Println("It while loop using for-loop")
	point := 1
	for point < 5 {
		fmt.Println(point)
		point++
	}
	fmt.Println("\n**********************")

	// infinite loop
	/* value := 1
	for {
		value++
	}
	*/

	// for-each : range
	greet := []string{"Hello", "Good", "Morning"}
	for i, g := range greet {
		fmt.Println(i, g)
	}
}
