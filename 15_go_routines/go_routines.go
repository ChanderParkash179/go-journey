package goroutines

import (
	"fmt"
	"time"
)

func GoRoutines() {

	go goRoutineFunc()

	nonGoRoutineFunc()
}

func goRoutineFunc() {
	fmt.Println("I am function 01")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("I am function 1 after 2 sec")
}

func nonGoRoutineFunc() {
	fmt.Println("I am function 02")
}
