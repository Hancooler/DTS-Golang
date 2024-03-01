package main

import (
	"fmt"
)

/*
func main(){
	var firstnumber *int // pointer type
	var secondNumber *int type

	_,_ = firstnumber, secondNumber
}
*/
/* pointer memory address
func main() {
	var firstNumber int = 10
	var secondNumber *int = &firstNumber

	fmt.Println("firstnumber(value) :", firstNumber)
	fmt.Println("firstnumber(memory address) :", firstNumber)

	fmt.Println("secondnumber(value) :", *secondNumber)
	fmt.Println("secondnumber(memory address) :", secondNumber)
}
*/

/*
	pointer changing value though pointer

	func main() {
		var firstperson string = "Jhon"
		var secondperson *string = &firstperson

		fmt.Println("firstperson (value) :", firstperson)
		fmt.Println("firstperson (memory address) :", &firstperson)
		fmt.Println("secondperson (value) :", *secondperson)
		fmt.Println("secondperson (memory address) :", secondperson)

		fmt.Println("\n", strings.Repeat("#", 30), "\n")

		*secondperson = "doe"

		fmt.Println("firstperson (value) :", firstperson)
		fmt.Println("firstperson (memory address) :", &firstperson)
		fmt.Println("secondperson (value) :", *secondperson)
		fmt.Println("secondperson (memory address) :", secondperson)
	}
*/
func main() {
	var a int = 10

	fmt.Println("before:", a)

	changevalue(&a)

	fmt.Println("after:", a)
}

func changevalue(number *int) {
	*number = 20
}
