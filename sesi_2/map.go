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

/* map with looping
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
*/

/* map deleting value
func main() {
	var person = map[string]string{
		"name":    "aril",
		"age":     "23",
		"address": "sudirman",
	}

	fmt.Println("before deleting:", person)

	delete(person, "age")

	fmt.Println("after deleting", person)

}
*/

/* map detecting a value
func main() {
	var person = map[string]string{
		"name":    "aril",
		"age":     "23",
		"address": "sudirman",
	}

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value doesn'nt exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value has been deleted")
	}

}
*/

func main() {
	var person = map[string]string{
		{"name": "aril", "age", "23"},
		{"name": "nanda", "age", "23"},
		{"name": "mail", "age", "23"},
	}

	for i, person := range person {
		fmt.Printf("index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}

}
