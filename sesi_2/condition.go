package main

import "fmt"

/* Condition (temporary variabel)
func main() {
	var currenyear = 2021

	if age := currenyear - 1998; age < 17 {
		fmt.Println("kamu belum boleh membuat kartu sim")

	} else {
		fmt.Println("kamu sudah boleh membuat kartu sim")
	}
}
*/

/* switch
func main() {

	var score = 7

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}
*/

/*	switch with relational operators
func main() {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}
}
*/

/* switch fallthrough keyword
func main() {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
		fallthrough
	case score < 5:
		fmt.Println("it is ok, but plwase study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you don't have a good score yet")
		}
	}
}
*/

func main() {
	var score = 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder")
			}
		}
	}
}
