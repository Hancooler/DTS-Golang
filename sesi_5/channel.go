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
/*channelwith unbuffer
func main() {
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("goroutine start sending data ")
		c <- 10
		fmt.Println("goroutine finish sending data ")
	}(c1)

	fmt.Println("goroutine sleep 2 second")
	time.Sleep(2 * time.Second)

	fmt.Println("main receive data")
	d := <-c1
	fmt.Println("main finish receive data", d)

	close(c1)
	time.Sleep(time.Second)
}
*/
/* goroutine with buffer
func main() {
	c1 := make(chan int)

	go func(c chan int) {
		for i := 1; i < 5; i++ {
			fmt.Println("goroutine start sending data\n ", i)
			c <- i
			fmt.Println("goroutine finish sending data\n ", i)

		}
		close(c)
	}(c1)

	fmt.Println("goroutine sleep 2 second")
	time.Sleep(2 * time.Second)

	for v := range c1 {
		fmt.Println("main receive data\n ", v)
	}
}
*/
/* channel with select
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "world"

	}()

	for i := 1; i < +2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
*/
