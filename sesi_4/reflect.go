package main

import (
	"fmt"
	"reflect"
)

/*
	identifying data type & value

	func main() {
		var number = 10
		var reflectValue = reflect.ValueOf(number)

		fmt.Println("tipe variabel:", reflectValue.Type())
		if reflectValue.Kind() == reflect.Int {
			fmt.Println("value variabel:", reflectValue.Int())

		}
	}
*/

/* Value Using Interface Method
func main() {
	var number = 10
	var reflectValue = reflect.ValueOf(number)
	//var nilai = reflectValue.Interface().(int)

	fmt.Println("tipe variabel:", reflectValue.Type())
	fmt.Println("nilai variabel:", reflectValue.Int())
}

*/
/*
func (s *student) setName(name string) {
	s.name = name
}
*/
type student struct {
	Name  string
	grade int
}

func main() {
	var s1 = &student{Name: "jhon", grade: 10}
	fmt.Println("nama:", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("setName")
	method.Call([]reflect.Value{
		reflect.ValueOf("jhon"),
	})

	fmt.Print("nama:", s1.Name)
}
