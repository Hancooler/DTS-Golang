package main

import (
	"fmt"
)

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

/* combining slicing and append
func main() {
	var fruits1 = []string{"apple", "banana", "manggo", "durian", "pineaplle"}
	fruits1 = append(fruits1[:3], "rambutan")
	fmt.Printf("%#v\n", fruits1)
}
*/

/*
	slice backing array

func main() {

		var fruits1 = []string{"apple", "manggo", "durian", "banana", "starfruit"}

		var fruits2 = fruits1[2:4]

		fruits2[0] = "rambutan"

		fmt.Println("fruit1 =>", fruits1)
		fmt.Println("fruit2 =>", fruits2)
	}
*/

/* slice cap function
func main() {
	var fruit1 = []string{"apple", "manggo", "durian", "banana"}

	fmt.Println("fruits1 cap:", cap(fruit1))
	fmt.Println("fruits1 len:", len(fruit1))

	fmt.Println(strings.Repeat("#", 20))

	var fruit2 = fruit1[0:3]

	fmt.Println("fruits2 cap:", cap(fruit2))
	fmt.Println("fruits2 len:", len(fruit2))

	fmt.Println(strings.Repeat("#", 20))

	var fruit3 = fruit1[1:]

	fmt.Println("fruits3 cap:", cap(fruit3))
	fmt.Println("fruits3 len:", len(fruit3))

	fmt.Println(strings.Repeat("#", 20))
}
*/

func main() {
	cars := []string{"ford", "honda", "audi", "rover"}
	newcars := []string{}

	newcars = append(newcars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newcars", newcars)
}
