package main

import "fmt"

func main() {
	fruits := [5]string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	for i, v := range fruits {
		fmt.Printf("Value[%s]\tv:Address[%p] fruits:IndexAddr[%p]\n", v, &v, &fruits[i])
	}

	for _, v := range fruits {
		v = v + "+v"
		fmt.Printf("Value[%s]\n", v)
	}
	fmt.Println(fruits)

	for i, v := range fruits {
		fruits[1] = "Jeck"
		fmt.Printf("i: %d Value[%s]\n", i, v)
	}
	fmt.Println(fruits)
}
