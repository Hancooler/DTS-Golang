package main

/* closure declare closure in variable
func main() {

	var evennumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(evennumbers(numbers...))
}
*/

/* closure IIFE (immediately invoked function expression)
func main() {
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("madam")

	fmt.Println(isPalindrome)
}
*/
/*
func main() {

	var studentlist = []string{"aril", "nanda", "nailo", "mikel", "marco"}

	var find = findtudent(studentlist)

	fmt.Println(find("marco"))
}
*/
/* closure as a return value
func findtudent(students []string) func(string) string {

	return func(name string) string {
		var student string
		var posisition int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(name) {
				student = v
				posisition = i
				break
			}
		}
		if student == "" {
			return "student not found"
		} else {
			return fmt.Sprintf("student %s found at position %d", student, posisition)

		}
	}
}
*/
/* callback
func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var find = findOddNumbers(numbers, func(numbers int) bool {
		return numbers%2 != 0

	})

	fmt.Println("total odd number", find)
}

func findOddNumbers(numbers []int, filter func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback := filter(v); callback {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
*/
