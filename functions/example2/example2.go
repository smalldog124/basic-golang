package main

import "fmt"

func main() {
	var n int

	// anonymous function and call it.
	func() {
		fmt.Println("Direct:", n)
	}()

	f := func() {
		fmt.Println("Variable:", n)
	}
	defer f()

	f()

	n = 20

	func(x int) {
		fmt.Println("Direct input parm:", x)
	}(80)
}
