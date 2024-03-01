package main

import (
	"fmt"
	"runtime"
)

/* goroutine
func main() {
	go goroutine()
}

func goroutine() {
	fmt.Println("hello goroutine")
}
*/

func main() {
	fmt.Println("main execution started")
	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No. of goroutine:", runtime.NumGoroutine())
	fmt.Println("main execution ended")
}

func firstProcess(index int) {
	fmt.Println("first process started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)

	}

	fmt.Println("first process ended")
}

func secondProcess(index int) {
	fmt.Println("second process started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)

	}

	fmt.Println("second process ended")
}
