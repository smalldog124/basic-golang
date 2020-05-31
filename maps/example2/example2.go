package main

import "fmt"

func main() {
	fruits := map[int]string{
		1: "Apple",
		2: "Orange",
		3: "Banana",
		4: "Grape",
		5: "Plum",
	}
	fmt.Printf("%+v\n", fruits)

	delete(fruits, 1)
	fmt.Println(fruits)

	orange, ok := fruits[2]
	fmt.Printf("value:%s \t have key: %t\n", orange, ok)

	_, ok = fruits[1]
	fmt.Printf("have key: %t\n", ok)
}
