package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))

	})

	fmt.Println("server is runnung on port 5000")

	http.HandleFunc("/employee", getEmployee)
	http.HandleFunc("/employee/create", createEmployee)

	http.ListenAndServe(":5000", nil)
}

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

func getEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// set header to return json
	w.Header().Set("Content-Type", "application/json")
	// set Status to 200 (ok)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// set header to return json
	w.Header().Set("Content-Type", "application/json")
	// set Status to 201 (created)
	w.WriteHeader(http.StatusCreated)

	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)

}
