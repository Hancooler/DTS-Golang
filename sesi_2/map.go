package main

import "fmt"

/* map
func main() {
	var person = map[string]string{
		"name":    "aril",
		"age":     "23",
		"address": "sudirman",
	}

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])
}
*/
func main() {
	var person = map[string]string{
		"name":    "aril",
		"age":     "23",
		"address": "sudirman",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}

}
