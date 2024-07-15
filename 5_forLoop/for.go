package main

import "fmt"

// for => Only construct in Go for looping

func main() {
	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinite loop
	// for{
	// 	println("1")
	// }

	// classic for loop

	// for i := 0; i < 3; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Range

	for i := range 3 {
		fmt.Println(i)
	}

}
