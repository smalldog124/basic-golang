package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Orange", "Banana", "Grape", "Plum"}

	inspectSlice(fruits)

	fruits2 := fruits[2:4]
	inspectSlice(fruits2)

	var food = make([]string, 3)
	copy(food, fruits)
	inspectSlice(food)

	greenFruits := []string{"Avocado", "GreenApple", "Mango"}
	fruits = append(fruits, greenFruits...)
	inspectSlice(fruits)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}
