package main

import "fmt"

func main() {
	// สร้างตัวแปร constants
	const vat = 0.07

	// สร้างตัวแปรแบบ Loag declaration
	var productName string
	var price float64
	var amount int

	fmt.Printf("vat: [%T] %f\n", vat, vat)
	fmt.Printf("productName: [%T] %s\n", productName, productName)
	fmt.Printf("price: [%T] %.2f\n", price, price)
	fmt.Printf("amount: [%T] %d\n", amount, amount)

	// สร้างตัวแปรแบบ Short declaration
	name := "smalldog"
	age := 20
	address := "shu ha ri company"
	isAdmin := false

	fmt.Printf("name: [%T] %s\n", name, name)
	fmt.Printf("age: [%T] %d\n", age, age)
	fmt.Printf("address: [%T] %s\n", address, address)
	fmt.Printf("isAdmin: [%T] %v\n", isAdmin, isAdmin)

	// กำหนดชนิดของข้อมูลแบบเจาะจง และเปลี่ยนชนิดข้อมูล
	userID := int16(9)
	fmt.Printf("userID: [%T] %d\n", userID, userID)
}
