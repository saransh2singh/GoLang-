package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch

	// i := 2
	// switch i {
	// case 1:
	// 	fmt.Println("1")
	// 	// No need of break statement
	// case 2:
	// 	fmt.Println("2")
	// case 3:
	// 	fmt.Println("3")
	// default:
	// 	fmt.Println("Other")
	// }

	// Multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")

	default:
		fmt.Println("It's workday")

	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("Other", t)
		}
	}
	whoAmI(1)
}
