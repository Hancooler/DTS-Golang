package main

/* error
func main() {
	var password string // type

	fmt.Scanln(&password)

	if valid, err := validpassword(password); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valid)

	}
}

func validpassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {

		return "", errors.New("password terlalu pendek")

	}
	return "valid", nil
}
*/
/* panic
func main() {
	var password string // type

	fmt.Scanln(&password)

	if valid, err := validpassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)

	}
}

func validpassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password terlalu pendek")

	}
	return "valid", nil
}
*/
/* recover
func main() {
	var password string // type

	fmt.Scanln(&password)

	if valid, err := validpassword(password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)

	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("No error")
	}
}

func validpassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password terlalu pendek")
	}
	return "valid", nil
}
*/
