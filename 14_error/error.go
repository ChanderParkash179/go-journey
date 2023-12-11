package error

import (
	"errors"
	"fmt"
)

func Errors() {

	// errors

	result, err := divide(4, 0)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Result: ", result)

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("you can't divide any value by 0")
	}

	return a / b, nil
}
