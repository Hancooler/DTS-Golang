package main

import "fmt"

// looping first way
/*
func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("angka", i)
	}
}
*/
// looping second way
/*
func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("angka", i)
	}
}
*/
// looping third way of looping
/*
func main() {
	var i = 0

	for {
		fmt.Println("angka", i)

		i++
		if i == 3 {
			break
		}
	}
}
*/
// break and continue
/*
func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}
		fmt.Println("angka", i)
	}
}
*/
// nested looping
/*
func main() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, "")
		}

		fmt.Println("\n")
	}
}
*/
// loopings label
func main() {
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan ke -", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
