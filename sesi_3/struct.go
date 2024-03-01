package main

/* struct
type Employee struct {
	name string
	age  int
	division string
}

func main(){

}
*/

/* struct giving value to struct
type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee Employee
	employee.name = "aril"

	employee.age = 23

	employee.division = "curriculum developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
}
*/

/* struct initializing struct
type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee = Employee{}
	employee.name = "aril"
	employee.age = 23
	employee.division = "curriculum developer"

	var employee2 = Employee{
		name:     "nanda",
		age:      23,
		division: "finance",
	}
	fmt.Println("employee1", employee)
	fmt.Println("employee2", employee2)
}
*/

/* pointer to a struct
type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee1 = Employee{name: "aril", age: 23, division: "curriculum developer"}

	var employee2 *Employee = &employee1 // pointer type

	fmt.Println("employee1 name:", employee1.name)
	fmt.Println("employee2 name:", employee1.name)

	fmt.Println(strings.Repeat("#", 20))

	employee2.name = "nanda"

	fmt.Println("employee1 name:", employee1.name)
	fmt.Println("employee2 name:", employee2.name)
}
*/

/*
embedded struct

	type person struct {
		name string
		age  int
	}

	type Employee struct {
		person   person
		division string
	}

	func main() {
		var employee1 = Employee{}

		employee1.person.name = "aril"
		employee1.person.age = 23
		employee1.division = "curriculum developer"

		fmt.Printf("%+v", employee1)
	}
*/

/* struct anonymus struct
type Person struct { // Capitalize "Person" to make it exported and match your usage
	name string
	age  int
}

func main() {
	var employee1 = struct {
		person   Person
		division string
	}{}
	employee1.person = Person{name: "Aril", age: 23} // Ensure "Person" matches the struct definition
	employee1.division = "Curriculum Developer"

	var employee2 = struct { // Correct the syntax for initializing `employee2`
		person   Person
		division string
	}{
		person:   Person{name: "Nanda", age: 23}, // Correct the syntax
		division: "Finance",
	}

	fmt.Printf("employee1: %+v\n", employee1)
	fmt.Printf("employee2: %+v\n", employee2)
}
*/

/* struct slice
type Person struct {
	name string
	age  int
}

func main() {
	var people = []Person{
		{name: "aril", age: 23},
		{name: "nanda", age: 23},
		{name: "mail", age: 23},
	}
	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}

*/

/* struct slice of anonymus struct
type Person struct {
	name string
	age  int
}

func main() {
	// Define employee as a slice of an anonymous struct containing Person and division
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "aril", age: 23}, division: "curriculum developer"},
		{person: Person{name: "nanda", age: 23}, division: "finance"},
		{person: Person{name: "mail", age: 23}, division: "marketing"},
	}
	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
}
*/
