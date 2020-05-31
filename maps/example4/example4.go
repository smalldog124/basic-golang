package main

import "fmt"

func main() {
	players := map[string]int{"Anna": 42,"Jacob": 21}
	fmt.Println(players)

	players["Anna"] ++
	fmt.Println(players)

	double(players,"Jacob")
	fmt.Println(players)
}

func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}