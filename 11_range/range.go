package main

import "fmt"

// iterating over data structures

func main() {

	// nums := []int{6, 7, 8}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// nums := []int{6, 7, 8}
	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println("Index : ", i, " Number : ", nums[i])
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	// m := map[string]string{"fname": "john", "lname": "doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// unicode code point of rune
	// starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
