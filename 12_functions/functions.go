package main

// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }

func processIt(fn func(a int) int) {
	fn(1)
}

func main() {
	// result := add(5, 5)
	// fmt.Println(result)

	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int {
		return 2
	}
	processIt(fn)
}
