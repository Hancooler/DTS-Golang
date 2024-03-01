package main

/* empty interface

func main() {
	var randomvalues interface{}

	_ = randomvalues

	randomvalues = "indonesia"

	randomvalues = 10

	randomvalues = true

	randomvalues = []string{"hello", "world"}

	fmt.Println("type of random: %T\n", randomvalues)
	fmt.Println("value of random: %v\n", randomvalues)
}
*/

/*emptyinterface assertion
func main() {
	var v interface{}

	v = 20
	if value, ok := v.(int); ok == true {
		v = value * 9

	}
}
*/

func main() {
	rs := []interface{}{1, "airel", true, 2, "nanda"}
	rm := map[string]interface{}{
		"name":   "aril",
		"status": true,
		"age":    23,
	}
	_, _ = rs, rm
}
