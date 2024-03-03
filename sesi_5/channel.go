package main

/* channel
func main() {
	c := make(chan string)

	go introduce("aril", c)
	go introduce("nanda", c)
	go introduce("nailo", c)

	msg1 := <-c
	msg2 := <-c
	msg3 := <-c
	fmt.Println(msg1, msg2, msg3)

	close(c)
}

func introduce(name string, c chan string) {
	result := fmt.Sprintf("hello %s", name)

	c <- result
}
*/

/* channelwith anonymous function
func main() {
	c := make(chan string)
	students := []string{"aril", "nanda", "nailo"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("hello %s", student)
			c <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}
*/

/* channel with direct function
func main() {
	c := make(chan string)
	students := []string{"aril", "nanda", "nailo"}

	for _, v := range students {
		go introduce(v, c)

	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("hello %s", student)
	c <- result
}
*/
