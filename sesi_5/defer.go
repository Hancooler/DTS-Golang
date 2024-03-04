package main

import (
	"fmt"
	"os"
)

/* defer
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
*/
/*
func main() {
	fmt.Println("hello")
	defer fmt.Println("world")
}

func calldeferfunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer startS")
}
*/

func main() {
	fmt.Println("hello")
	defer fmt.Println("world")
	os.Exit(1)
}
