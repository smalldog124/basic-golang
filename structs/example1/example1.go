package main

import "fmt"

type user struct{
  name string
  age int
  weigth float32
  high int
}

func main() {
	u := user{}
	fmt.Printf("%+v\n",u)

	jone := user{
		name: "jone",
		age: 25,
		weigth: 70.00,
		high: 183,
	}
	fmt.Printf("name: %s\n",jone.name)
	fmt.Printf("age: %d\n",jone.age)
	fmt.Printf("weigth: %f\n",jone.weigth)
	fmt.Printf("high: %d\n",jone.high)


	product := struct{
		id int
		name string
		princr float64
	}{
		1,
		"Macbook Pro",
		59900,
	}
	fmt.Printf("%+v",product)
}