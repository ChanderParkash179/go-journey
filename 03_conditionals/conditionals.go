package conditionals

import "fmt"

func Conditionals() {

	age, day := 23, 6

	ifElse(age)
	switchCase(age)

	multiCasesValueSwitch(day)
}

func ifElse(age int) {
	if age > 18 {
		fmt.Println("you're eligible!")
	} else if age == 18 {
		fmt.Println("you're eligible completely!")
	} else {
		fmt.Println("you're not eligible!")
	}
}

func switchCase(age int) {
	switch age {
	case 19:
		fmt.Println("you're eligible!")
	case 18:
		fmt.Println("you're eligible completely!")
	default:
		fmt.Println("you're not eligible!")
	}
}

func multiCasesValueSwitch(day int) {
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("its weekday")
	case 6, 7:
		fmt.Println("its weekend")
	default:
		fmt.Println("invalid day provided")
	}
}
