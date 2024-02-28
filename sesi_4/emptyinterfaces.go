package main

import "fmt"

func main() {
	var randomvalues interface{}

	_ = randomvalues

	randomvalues = "indonesia"

	randomvalues = 10

	randomvalues = true

	randomvalues = []string{"hello", "world"}

	fmt.Println("type of random: %T\n", randomvalues)
	fmt.Println("value of random: %v\n", randomvalues)
}
