package main

import (
	"fmt"
	"maps"
)

// maps => hash,object,dict

func main() {
	// Creating map
	// m := make(map[string]string)
	// setting element
	// m["name"] = "golang"
	// m["area"] = "backend"
	// get an element
	// fmt.Println(m["name"], m["area"])

	// if key doesn't exists in the map then it returns zero value
	// fmt.Println(m["phome"])

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["age"])
	// fmt.Println(m["phone"]) // 0
	// fmt.Println(len(m))
	// delete(m, "price")
	// fmt.Println(m)
	// clear(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}
	// // fmt.Println(m)
	// v, ok := m["price"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("All ok")
	// } else {
	// 	fmt.Println("Not ok")
	// }

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}
	fmt.Println(maps.Equal(m1, m2))

}
