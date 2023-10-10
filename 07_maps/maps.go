package maps

import "fmt"

func Maps() {
	fmt.Println("maps : start")

	var mp = map[string]int{
		"ali":   23,
		"ahmed": 32,
	}

	fmt.Println(mp["ali"])
	fmt.Println(mp["ahmed"])

	fmt.Println("maps : end")
}
