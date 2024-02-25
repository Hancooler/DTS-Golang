package main

import "fmt"

func main() {
	var currenyear = 2021

	if age := currenyear - 1998; age < 17 {
		fmt.Println("kamu belum boleh membuat kartu sim")

	} else {
		fmt.Println("kamu sudah boleh membuat kartu sim")
	}
}
