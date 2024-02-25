package main

import "fmt"

/* variable with tipe data
func main() {
	var name string = "hans"
	var age int = 23

	fmt.Println("ini namanaya ==>", name)
	fmt.Println("ini umurnya ==>", age)
}
*/
// variable without tipe data
func main() {
	//var name = "ariel"
	//var age int = 23

	//fmt.Println("ini namanaya ==>", name)
	//fmt.Println("ini umurnya ==>", age)

	//fmt.Printf("%T,%T", name, age)

	// multiple variable declarations

	//	var std1, std2, std3 string = "satu", "dua", "tiga"
	//	var first, second, third int
	//	first, second, third = 1, 2, 3

	//	fmt.Println(std1, std2, std3)
	//	fmt.Println(first, second, third)

	// underscore variabel
	//var firstvariable string

	//var name, age, address = "aril", 23, "jalan sudirman"

	//_, _, _, _ = firstvariable, name, age, address

	// fmt.print
	//first, second := 1, "2"

	//fmt.Printf("tpdata first %T\n", first)
	//fmt.Printf("tpdata Second %T\n", second)

	var nama = "aril"

	var age = 23

	var address = "jalan sudirman"

	fmt.Printf("halo nama saya %s, umur saya %d, aku tinggal $s", nama, age, address)
}
