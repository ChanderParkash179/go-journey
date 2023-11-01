package typeinference

import "fmt"

func TypeInference() {

	fmt.Println("Type Inference!")

	var name, age, hasLicense = "chander parkash", 24, true

	fmt.Printf("name is type of %T, age is type of %T, hasLicense is type of %T\n", name, age, hasLicense)
}
