package main

/* slice make function
func main() {
	var fruits = make([]string, 3)

	_ = fruits

	fmt.Printf("%#v", fruits)
}
*/

/* slice append function
func main() {
	var fruits = make([]string, 3)

	fruits = append(fruits, "apple", "banana", "manggo")

	fmt.Printf("%#v", fruits)
}
*/

/* slice append function with ellipsis
func main() {
	var fruits1 = []string{"apple", "banana", "manggo"}

	var fruits2 = []string{"string", "pineapple", "starfruits"}

	fruits1 = append(fruits1, fruits2...)

	fmt.Printf("%#v", fruits1)
}
*/

/* slice copy function
func main() {
	var fruits1 = []string{"apple", "banana", "manggo"}

	var fruits2 = []string{"string", "pineapple", "starfruits"}

	nn := copy(fruits1, fruits2)

	fmt.Println("fruits1 =>", fruits1)
	fmt.Println("fruits2 =>", fruits2)
	fmt.Println("copied elements =>", nn)
}
*/

/* slice slicing
func main() {
	var fruits1 = []string{"apple", "banana", "manggo", "durian", "pineaplle"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("%#v\n", fruits3)

	var fruits4 = fruits1[:3]
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits1[:]
	fmt.Printf("%#v\n", fruits5)
}
*/
