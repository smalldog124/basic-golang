package example2

import "fmt"

func FizzBuzz(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	}
	if num%5 == 0 {
		return "Buzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", num)
}
