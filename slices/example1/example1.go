package main

import "fmt"

func main() {
	// ประกาศตัวแปร fruits ที่มีขนาด 5
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// เพิ่มข้อมูลเข้าไปอีกที่เกิน length ที่ประกาศไว้
	fruits[5] = "Runtime error"

	// Error: panic: runtime error: index out of range

	fmt.Println(fruits)
}
