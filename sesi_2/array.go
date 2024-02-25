package main

import (
	"fmt"
)

/* array
func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"arol", "nada", "milo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)

}
*/
// array modify element through index
/*func main() {
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "manggo"

	fmt.Printf("%#v\n", fruits)
}
*/

/* array loop through elements
func main() {
	var fruits = [3]string{"apel", "pisang", "mangga"}
	// cara pertama
	for i, v := range fruits {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	//cara kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index: %d, value: %s\n", i, fruits[i])
	}
}
*/

func main() {
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
