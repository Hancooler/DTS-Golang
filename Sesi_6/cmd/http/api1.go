package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

var employees = []Employee{
	{ID: 1, Name: "aril", Age: 23, Division: "curriculum developer"},
	{ID: 2, Name: "nanda", Age: 23, Division: "finance"},
	{ID: 3, Name: "mail", Age: 23, Division: "marketing"},
}

var PORT = ":5000"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee) // Ini diperbaiki dari http.Handle ke http.HandleFunc

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid request method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		err := r.ParseForm() // Tambahkan baris ini untuk parse form data
		if err != nil {
			http.Error(w, "Error parsing the form", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid request method", http.StatusBadRequest)
}
