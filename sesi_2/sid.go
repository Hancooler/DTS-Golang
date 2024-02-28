// string in depth
package main

import (
	"fmt"
)

/* looping over string (byte by byte)
func main() {
	city := "jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}
}
*/
/*
func main() {
	var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(city))
}
*/

// loop over string rune by rune
/*
func main() {
	str1 := "makan"
	str2 := "manca"

	fmt.Printf("str1 byte lenght => %d\n", len(str1))
	fmt.Printf("str2 byte lenght => %d\n", len(str2))
}
*/
func main() {
	str := "manca"

	for i, s := range str {
		fmt.Printf("index =>%d, dtring => %s\n", i, string(s))
	}
}
