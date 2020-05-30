package main

import "fmt"

func main() {
	var numbers [4]int
	fmt.Println(numbers)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 200
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("i: [%d] number: %d\n", i, numbers[i])
	}

	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i, friend := range friends {
		fmt.Println(i, friend)
	}
}
