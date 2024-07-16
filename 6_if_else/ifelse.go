package main

import "fmt"

func main() {
	// age := 18
	// if age >= 19 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// age := 33
	// if age >= 18 {
	// 	fmt.Println("Person is an adult", age)
	// } else if age >= 12 {
	// 	fmt.Println("Person is teenager")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" || hasPermissions {
	// 	fmt.Println("Yes")
	// }

	// if role == "admin" && hasPermissions {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("No")
	// }

	// We can declare a variable inside if
	if age := 19; age >= 18 {
		fmt.Println("Adult age: ", age)
	} else if age >= 12 {
		fmt.Println("Person is a teenager", age)
	}

	// Go doesn't have ternary operator, you'll have to use if else

}
