package main

import (
	"fmt"
)

/* function
func main() {
	greet("aril", 23)
}

func greet(name string, age int8) {
	fmt.Printf("hello saya %s and  saya %d tahun", name, age)
}
*/

/*
func main() {
	greet("arin", "jalan sudirman")

}

func greet(name, address string) {
	fmt.Println("hello there my name is", name)
	fmt.Println("i live in", address)
}
*/

/* Function Return
func main() {
	var names = []string{"aril", "jordan"}
	var printmsg = greet("hei", names)

	fmt.Println(printmsg)
}

func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}
*/

/* funtcion returning multiple values
func main() {
	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Println("area:", area)
	fmt.Println("circumference:", circumference)
}

func calculate(d float64) (float64, float64) {
	//menghitung luas
	var area float64 = math.Pi * d
	//menghitung keliling

	var circumference = math.Pi * d

	return area, circumference
}
*/

/* function predefined return value
func main() {
	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Println("area:", area)
	fmt.Println("circumference:", circumference)

}

func calculate(d float64) (area float64, circumference float64) {
	//menghitung luas
	area = math.Pi * math.Pow(d/2, 2)
	//menghitung keliling

	circumference = math.Pi * d

	return
}
*/

func main() {
	studenlist := print("aril", "nanda", "nailo", "mikel", "marco")

	fmt.Printf("%v", studenlist)

}
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)

	}
	return result
}
