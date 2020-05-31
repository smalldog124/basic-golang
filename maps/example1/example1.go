package main

import "fmt"

func main() {
	number := make(map[string]int)
	number["one"] = 1
	number["two"] = 2

	fmt.Println("number:", number)

	namePrefix := map[string]string{"Mr": "นาย", "Ms": "นางสาว"}
	fmt.Println("namePrefix:", namePrefix)

	fmt.Println(namePrefix["Ms"])
}
