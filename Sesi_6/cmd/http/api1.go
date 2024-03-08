package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{

	{ID: 1, Name: "aril", Age: 23, Division: "curriculum developer"},
	{ID: 2, Name: "nanda", Age: 23, Division: "finance"},
	{ID: 3, Name: "mail", Age: 23, Division: "marketing"},
}

var PORT = ":5000"

func main() {
	http.HandleFunc("/employees", getEmployees)
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
